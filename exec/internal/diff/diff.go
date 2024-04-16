// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package diff

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// 一对是diff中同时跟踪x和y两侧的两个值。通常是一对行索引。
type pair struct{ x, y int }

// Diff returns an anchored diff of the two texts old and new
// in the “unified diff” format. If old and new are identical,
// Diff returns a nil slice (no output).
//
// Unix diff implementations typically look for a diff with
// the smallest number of lines inserted and removed,
// which can in the worst case take time quadratic in the
// number of lines in the texts. As a result, many implementations
// either can be made to run for a long time or cut off the search
// after a predetermined amount of work.
//
// In contrast, this implementation looks for a diff with the
// smallest number of “unique” lines inserted and removed,
// where unique means a line that appears just once in both old and new.
// We call this an “anchored diff” because the unique lines anchor
// the chosen matching regions. An anchored diff is usually clearer
// than a standard diff, because the algorithm does not try to
// reuse unrelated blank lines or closing braces.
// The algorithm also guarantees to run in O(n log n) time
// instead of the standard O(n²) time.
//
// Some systems call this approach a “patience diff,” named for
// the “patience sorting” algorithm, itself named for a solitaire card game.
// We avoid that name for two reasons. First, the name has been used
// for a few different variants of the algorithm, so it is imprecise.
// Second, the name is frequently interpreted as meaning that you have
// to wait longer (to be patient) for the diff, meaning that it is a slower algorithm,
// when in fact the algorithm is faster than the standard one.

// ff:
// new:
// newName:
// old:
// oldName:
func Diff(oldName string, old []byte, newName string, new []byte) []byte {
	if bytes.Equal(old, new) {
		return nil
	}
	x := lines(old)
	y := lines(new)

	// Print diff header.
	var out bytes.Buffer
	fmt.Fprintf(&out, "diff %s %s\n", oldName, newName)
	fmt.Fprintf(&out, "--- %s\n", oldName)
	fmt.Fprintf(&out, "+++ %s\n", newName)

// 遍历待考虑的匹配项，
// 将每个匹配项扩展至包含周围的行，
// 然后打印差异块。
// 为避免循环外的设置/清理情况，
// tgs 在匹配项序列中返回一个前导的 {0,0} 和尾部的 {len(x), len(y)} 对。
	var (
		done  pair     // 已打印至 x[:done.x] 和 y[:done.y]
		chunk pair     // 当前块的起始行
		count pair     // 当前块中每侧的行数
		ctext []string // lines for current chunk
	)
	for _, m := range tgs(x, y) {
		if m.x < done.x {
			// 已经处理了从先前匹配处向前扫描的情况
			continue
		}

// 尽可能地扩展匹配行，
// 确保 x[start.x:end.x] 与 y[start.y:end.y] 相等。
// 注意在第一次（或最后一次）迭代中，我们可能（或肯定）会遇到空匹配：start.x==end.x 且 start.y==end.y。
		start := m
		for start.x > done.x && start.y > done.y && x[start.x-1] == y[start.y-1] {
			start.x--
			start.y--
		}
		end := m
		for end.x < len(x) && end.y < len(y) && x[end.x] == y[end.y] {
			end.x++
			end.y++
		}

// 将起始点之前的不匹配行输出到此块中。
// （在第一个哨兵迭代时无影响，此时 start = {0,0}。）
		for _, s := range x[done.x:start.x] {
			ctext = append(ctext, "-"+s)
			count.x++
		}
		for _, s := range y[done.y:start.y] {
			ctext = append(ctext, "+"+s)
			count.y++
		}

// 如果我们没有到达EOF（文件结束符）且共同行数过少，
// 则该块应包含所有共同行并继续。
		const C = 3 // 上下文行数
		if (end.x < len(x) || end.y < len(y)) &&
			(end.x-start.x < C || (len(ctext) > 0 && end.x-start.x < 2*C)) {
			for _, s := range x[start.x:end.x] {
				ctext = append(ctext, " "+s)
				count.x++
				count.y++
			}
			done = end
			continue
		}

		// 以包含上下文的通用行结束数据块
		if len(ctext) > 0 {
			n := end.x - start.x
			if n > C {
				n = C
			}
			for _, s := range x[start.x : start.x+n] {
				ctext = append(ctext, " "+s)
				count.x++
				count.y++
			}
			done = pair{start.x + n, start.y + n}

// 格式化并输出代码块。
// 将行号转换为1索引制。
// 特殊情况：空文件显示为0,0而非1,0。
			if count.x > 0 {
				chunk.x++
			}
			if count.y > 0 {
				chunk.y++
			}
			fmt.Fprintf(&out, "@@ -%d,%d +%d,%d @@\n", chunk.x, count.x, chunk.y, count.y)
			for _, s := range ctext {
				out.WriteString(s)
			}
			count.x = 0
			count.y = 0
			ctext = ctext[:0]
		}

		// 如果我们到达了EOF，就完成了。
		if end.x >= len(x) && end.y >= len(y) {
			break
		}

		// 否则开始一个新的分块
		chunk = pair{end.x - C, end.y - C}
		for _, s := range x[chunk.x:end.x] {
			ctext = append(ctext, " "+s)
			count.x++
			count.y++
		}
		done = end
	}

	return out.Bytes()
}

// lines returns the lines in the file x, including newlines.
// If the file does not end in a newline, one is supplied
// along with a warning about the missing newline.
func lines(x []byte) []string {
	l := strings.SplitAfter(string(x), "\n")
	if l[len(l)-1] == "" {
		l = l[:len(l)-1]
	} else {
// 将最后一行视为附带有关缺失换行符的消息，
// 使用与 BSD/GNU diff 相同的文本（包括前导反斜杠）。
		l[len(l)-1] += "\n\\ No newline at end of file\n"
	}
	return l
}

// tgs returns the pairs of indexes of the longest common subsequence
// of unique lines in x and y, where a unique line is one that appears
// once in x and once in y.
//
// The longest common subsequence algorithm is as described in
// Thomas G. Szymanski, “A Special Case of the Maximal Common
// Subsequence Problem,” Princeton TR #170 (January 1975),
// available at https://research.swtch.com/tgs170.pdf.
func tgs(x, y []string) []pair {
// 统计字符串a和b中各字符串出现的次数。
// 我们仅关注每个字符串在a、b中出现0次、1次或多于1次的情况，分别计为0、-1、-2（对于x侧）以及0、-4、-8（对于y侧）。
// 现在使用负数，以便后续区分正向行号。
	m := make(map[string]int)
	for _, s := range x {
		if c := m[s]; c > -2 {
			m[s] = c - 1
		}
	}
	for _, s := range y {
		if c := m[s]; c > -8 {
			m[s] = c - 4
		}
	}

// 现在，可以通过 m[s] = -1+-4 来识别唯一的字符串。
// 
// 收集 x 和 y 中这些字符串的索引，构建如下：
//  xi[i] = x 中唯一字符串递增的索引。
//  yi[i] = y 中唯一字符串递增的索引。
// inv[i] = 索引 j，满足 x[xi[i]] = y[yi[j]]。
	var xi, yi, inv []int
	for i, s := range y {
		if m[s] == -1+-4 {
			m[s] = len(yi)
			yi = append(yi, i)
		}
	}
	for i, s := range x {
		if j, ok := m[s]; ok && j >= 0 {
			xi = append(xi, i)
			inv = append(inv, j)
		}
	}

// 应用Szymanski论文中的算法A。
// 用该文中的术语表示，A = J = inv，B = [0, n)。
// 我们在返回的序列中添加哨兵对{0,0}和{len(x),len(y)}，以辅助处理循环。
	J := inv
	n := len(xi)
	T := make([]int, n)
	L := make([]int, n)
	for i := range T {
		T[i] = n + 1
	}
	for i := 0; i < n; i++ {
		k := sort.Search(n, func(k int) bool {
			return T[k] >= J[i]
		})
		T[k] = J[i]
		L[i] = k + 1
	}
	k := 0
	for _, v := range L {
		if k < v {
			k = v
		}
	}
	seq := make([]pair, 2+k)
	seq[1+k] = pair{len(x), len(y)} // sentinel at end
	lastj := n
	for i := n - 1; i >= 0; i-- {
		if L[i] == k && J[i] < lastj {
			seq[k] = pair{xi[i], yi[J[i]]}
			k--
		}
	}
	seq[0] = pair{0, 0} // sentinel at start
	return seq
}
