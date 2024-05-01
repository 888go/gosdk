//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package strconv_test

import (
	"bufio"
	"fmt"
	"github.com/888go/gosdk/strconv"
	"os"
	"strings"
	"testing"
)

func pow2(i int) float64 {
	switch {
	case i < 0:
		return 1 / pow2(-i)
	case i == 0:
		return 1
	case i == 1:
		return 2
	}
	return pow2(i/2) * pow2(i-i/2)
}

// 这是一个 strconv.ParseFloat(x, 64) 的包装器。它自身处理 dddddp+ddd（二进制指数）部分，其余的传递给 strconv.ParseFloat。
// md5:da10d8aa57bdf05c
func myatof64(s string) (f float64, ok bool) {
	if mant, exp, ok := strings.Cut(s, "p"); ok {
		n, err := strconv.ParseInt(mant, 10, 64)
		if err != nil {
			return 0, false
		}
		e, err1 := strconv.Atoi(exp)
		if err1 != nil {
			println("bad e", exp)
			return 0, false
		}
		v := float64(n)
// 我们期望 v*pow2(e) 能够适应一个 float64 的范围，
// 但 pow2(e) 单独计算时可能超出了这个范围。因此需要格外小心。
// md5:f5e925857d3ba2e6
		if e <= -1000 {
			v *= pow2(-1000)
			e += 1000
			for e < 0 {
				v /= 2
				e++
			}
			return v, true
		}
		if e >= 1000 {
			v *= pow2(1000)
			e -= 1000
			for e > 0 {
				v *= 2
				e--
			}
			return v, true
		}
		return v * pow2(e), true
	}
	f1, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	return f1, true
}

// 包装了strconv.ParseFloat(x, 32)。处理dddddp+ddd（二进制指数）本身，其余的传递给strconv.ParseFloat。
// md5:a03fed9d2a4eb26f
func myatof32(s string) (f float32, ok bool) {
	if mant, exp, ok := strings.Cut(s, "p"); ok {
		n, err := strconv.Atoi(mant)
		if err != nil {
			println("bad n", mant)
			return 0, false
		}
		e, err1 := strconv.Atoi(exp)
		if err1 != nil {
			println("bad p", exp)
			return 0, false
		}
		return float32(float64(n) * pow2(e)), true
	}
	f64, err1 := strconv.ParseFloat(s, 32)
	f1 := float32(f64)
	if err1 != nil {
		return 0, false
	}
	return f1, true
}

func TestFp(t *testing.T) {
	f, err := os.Open("testdata/testfp.txt")
	if err != nil {
		t.Fatal("testfp: open testdata/testfp.txt:", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for lineno := 1; s.Scan(); lineno++ {
		line := s.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		a := strings.Split(line, " ")
		if len(a) != 4 {
			t.Error("testdata/testfp.txt:", lineno, ": wrong field count")
			continue
		}
		var s string
		var v float64
		switch a[0] {
		case "float64":
			var ok bool
			v, ok = myatof64(a[2])
			if !ok {
				t.Error("testdata/testfp.txt:", lineno, ": cannot atof64 ", a[2])
				continue
			}
			s = fmt.Sprintf(a[1], v)
		case "float32":
			v1, ok := myatof32(a[2])
			if !ok {
				t.Error("testdata/testfp.txt:", lineno, ": cannot atof32 ", a[2])
				continue
			}
			s = fmt.Sprintf(a[1], v1)
			v = float64(v1)
		}
		if s != a[3] {
			t.Error("testdata/testfp.txt:", lineno, ": ", a[0], " ", a[1], " ", a[2], " (", v, ") ",
				"want ", a[3], " got ", s)
		}
	}
	if s.Err() != nil {
		t.Fatal("testfp: read testdata/testfp.txt: ", s.Err())
	}
}
