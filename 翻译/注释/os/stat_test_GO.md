
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// testStatAndLstat verifies that all os.Stat, os.Lstat os.File.Stat and os.Readdir work.
<原文结束>

# <翻译开始>
// testStatAndLstat 验证所有 os.Stat、os.Lstat、os.File.Stat 和 os.Readdir 的功能。
# <翻译结束>


<原文开始>
// test fs.FileInfo returned by os.Readdir
<原文结束>

# <翻译开始>
// 测试由 os.Readdir 返回的 fs.FileInfo
# <翻译结束>


<原文开始>
// skip os.Readdir test of directories with slash at the end
<原文结束>

# <翻译开始>
// 跳过以斜杠结尾的目录的os.Readdir测试
# <翻译结束>


<原文开始>
// testIsDir verifies that fi refers to directory.
<原文结束>

# <翻译开始>
// testIsDir 验证fi是否指的是目录。
# <翻译结束>


<原文开始>
// testIsSymlink verifies that fi refers to symlink.
<原文结束>

# <翻译开始>
// testIsSymlink 验证fi是否引用符号链接。
# <翻译结束>


<原文开始>
// testIsFile verifies that fi refers to file.
<原文结束>

# <翻译开始>
// testIsFile 验证 fi 指向的是文件。
# <翻译结束>


<原文开始>
// see issue 27225 for details
<原文结束>

# <翻译开始>
// 查看问题 27225 以获取详细信息
# <翻译结束>

