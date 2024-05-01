
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2011 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
//
//func (r *Replacer) Replacer() any {
//	r.once.Do(r.buildOnce)
//	return r.r
//}
//
//func (r *Replacer) PrintTrie() string {
//	r.once.Do(r.buildOnce)
//	gen := r.r.(*genericReplacer)
//	return gen.printNode(&gen.root, 0)
//}
//
//func (r *genericReplacer) printNode(t *trieNode, depth int) (s string) {
//	if t.priority > 0 {
//		s += "+"
//	} else {
//		s += "-"
//	}
//	s += "\n"
//
//	if t.prefix != "" {
//		s += Repeat(".", depth) + t.prefix
//		s += r.printNode(t.next, depth+len(t.prefix))
//	} else if t.table != nil {
//		for b, m := range r.mapping {
//			if int(m) != r.tableSize && t.table[m] != nil {
//				s += Repeat(".", depth) + string([]byte{byte(b)})
//				s += r.printNode(t.table[m], depth+1)
//			}
//		}
//	}
//	return
//}
//
//func StringFind(pattern, text string) int {
//	return makeStringFinder(pattern).next(text)
//}
//
//func DumpTables(pattern string) ([]int, []int) {
//	finder := makeStringFinder(pattern)
//	return finder.badCharSkip[:], finder.goodSuffixSkip
//}
<原文结束>

# <翻译开始>
// ```
// Replacer方法返回一个替换器实例
// func (r *Replacer) Replacer() any {
//     r.once.Do(r.buildOnce)
//     return r.r
// }
// 
// PrintTrie方法打印替换器的字典树结构
// func (r *Replacer) PrintTrie() string {
//     r.once.Do(r.buildOnce)
//     gen := r.r.(*genericReplacer)
//     return gen.printNode(&gen.root, 0)
// }
// 
// genericReplacer的printNode方法递归打印字典树节点
// func (r *genericReplacer) printNode(t *trieNode, depth int) string {
//     s := ""
//     if t.priority > 0 {
//         s += "+"
//     } else {
//         s += "-"
//     }
//     s += "\n"
// 
//     if t.prefix != "" {
//         s += Repeat(".", depth) + t.prefix
//         s += r.printNode(t.next, depth+len(t.prefix))
//     } else if t.table != nil {
//         for b, m := range r.mapping {
//             if int(m) != r.tableSize && t.table[m] != nil {
//                 s += Repeat(".", depth) + string([]byte{byte(b)})
//                 s += r.printNode(t.table[m], depth+1)
//             }
//         }
//     }
//     return s
// }
// 
// StringFind函数在文本中查找匹配给定模式的位置
// func StringFind(pattern, text string) int {
//     return makeStringFinder(pattern).next(text)
// }
// 
// DumpTables函数返回模式的坏字符跳过和好后缀跳过数组
// func DumpTables(pattern string) ([]int, []int) {
//     finder := makeStringFinder(pattern)
//     return finder.badCharSkip[:], finder.goodSuffixSkip
// }
// ```
// 
// 这段代码是Go语言实现的一些字符串处理函数，包括创建替换器、打印替换器的字典树、查找模式在文本中的位置以及获取模式的坏字符跳过和好后缀跳过信息。
// md5:282dcfad4576bd45
# <翻译结束>

