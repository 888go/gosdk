
<原文开始>
// This file tests the situation where byte operations are checking
// data very near to a page boundary. We want to make sure those
// operations do not read across the boundary and cause a page
// fault where they shouldn't.
<原文结束>

# <翻译开始>
// 本文件测试了字节操作检查靠近页面边界数据的情况。
// 我们要确保这些操作不会读取跨越边界并引发不应有的页面错误。
// md5:75426777058c008d
# <翻译结束>


<原文开始>
// These tests run only on linux. The code being tested is
// not OS-specific, so it does not need to be tested on all
// operating systems.
<原文结束>

# <翻译开始>
// 这些测试仅在Linux上运行。被测试的代码不依赖特定操作系统，因此不需要在所有操作系统上进行测试。
// md5:88ba7e84796d0671
# <翻译结束>


<原文开始>
// dangerousSlice returns a slice which is immediately
// preceded and followed by a faulting page.
<原文结束>

# <翻译开始>
// dangerousSlice 返回一个切片，它立即由一个出错的页面前后跟随。
// md5:81173505fbc4ca2e
# <翻译结束>


<原文开始>
// Only worry about when we're near the end of a page.
<原文结束>

# <翻译开始>
// 只有在接近页面结束时才关心这个。. md5:b8599e6bc302cde8
# <翻译结束>


<原文开始>
// difference is only found on the last byte
<原文结束>

# <翻译开始>
// 差异只在最后一个字节上发现. md5:15407f2f4c5ab9e5
# <翻译结束>


<原文开始>
// Test differing alignments and sizes of q which always end on a page boundary.
<原文结束>

# <翻译开始>
// 测试在页边界结束时，不同对齐方式和大小的q。. md5:257221af700ea72a
# <翻译结束>

