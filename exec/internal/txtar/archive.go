// 版权所有 ? 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

// Package txtar 实现了一个简单的基于文本的文件归档格式。
//
// 该格式的目标如下：
//
//   - 简单到可以手动创建和编辑。
//   - 能够存储描述 go 命令测试用例的文本文件树。
//   - 在 git 历史记录和代码审查中易于进行 diff 比较。
//
// 非目标包括成为一个完全通用的归档格式、存储二进制数据、存储文件模式、存储符号链接等特殊文件等。
//
// # Txtar 格式
//
// 一个 txtar 归档由零个或多个注释行以及随后的一系列文件条目组成。每个文件条目以形如 "-- FILENAME --" 的文件标记行开始，后面跟随零个或多个构成文件数据的文件内容行。注释或文件内容在下一个文件标记行处结束。文件标记行必须以三字节序列 "-- " 开头，并以三字节序列 " --" 结尾，但其中包含的文件名两侧可以有额外的空白字符，这些空白字符都会被去除。
//
// 如果 txtar 文件在最后一行缺少尾部换行符，解析器应视其末尾存在换行符。
//
// txtar 归档中不存在任何可能的语法错误。
package txtar

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Archive 是一组文件的集合。
type Archive struct {
	Comment []byte
	Files   []File
}

// File 是档案中的单个文件。
type File struct {
	Name string // 文件名（"foo/bar.txt"）
	Data []byte // text content of file
}

// Format 函数返回 Archive 的序列化形式。
// 假定 Archive 数据结构是格式良好的：
// 其中 a.Comment 及所有 a.File[i].Data 不包含文件标记行，
// 并且所有 a.File[i].Name 非空。

// ff:
// a:
func Format(a *Archive) []byte {
	var buf bytes.Buffer
	buf.Write(fixNL(a.Comment))
	for _, f := range a.Files {
		fmt.Fprintf(&buf, "-- %s --\n", f.Name)
		buf.Write(fixNL(f.Data))
	}
	return buf.Bytes()
}

// ParseFile 以档案格式解析指定名称的文件。

// ff:
// file:
func ParseFile(file string) (*Archive, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(data), nil
}

// Parse解析Archive的序列化形式。
// 返回的Archive持有数据切片。

// ff:
// data:
func Parse(data []byte) *Archive {
	a := new(Archive)
	var name string
	a.Comment, name, data = findFileMarker(data)
	for name != "" {
		f := File{name, nil}
		f.Data, name, data = findFileMarker(data)
		a.Files = append(a.Files, f)
	}
	return a
}

var (
	newlineMarker = []byte("\n-- ")
	marker        = []byte("-- ")
	markerEnd     = []byte(" --")
)

// findFileMarker 在 data 中查找下一个文件标记，
// 提取文件名，并返回标记前的数据、文件名以及标记后的数据。
// 若不存在下一个标记，findFileMarker 返回 before = fixNL(data)，name = ""，after = nil。
func findFileMarker(data []byte) (before []byte, name string, after []byte) {
	var i int
	for {
		if name, after = isMarker(data[i:]); name != "" {
			return data[:i], name, after
		}
		j := bytes.Index(data[i:], newlineMarker)
		if j < 0 {
			return fixNL(data), "", nil
		}
		i += j + 1 // 位于新可能标记的起始处
	}
}

// isMarker 检查 data 是否以文件标记行开头。
// 如果是，它将返回该行中的名称以及该行之后的数据。
// 否则返回 name == "" 以及未指定的 after。
func isMarker(data []byte) (name string, after []byte) {
	if !bytes.HasPrefix(data, marker) {
		return "", nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		data, after = data[:i], data[i+1:]
	}
	if !(bytes.HasSuffix(data, markerEnd) && len(data) >= len(marker)+len(markerEnd)) {
		return "", nil
	}
	return strings.TrimSpace(string(data[len(marker) : len(data)-len(markerEnd)])), after
}

// 如果data为空或以\n结尾，fixNL函数返回data。
// 否则，fixNL函数返回一个新的切片，该切片由data加上末尾的\n组成。
func fixNL(data []byte) []byte {
	if len(data) == 0 || data[len(data)-1] == '\n' {
		return data
	}
	d := make([]byte, len(data)+1)
	copy(d, data)
	d[len(data)] = '\n'
	return d
}
