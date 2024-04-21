
<原文开始>
// Readdir reads the contents of the directory associated with file and
// returns a slice of up to n FileInfo values, as would be returned
// by Lstat, in directory order. Subsequent calls on the same file will yield
// further FileInfos.
//
// If n > 0, Readdir returns at most n FileInfo structures. In this case, if
// Readdir returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdir returns all the FileInfo from the directory in
// a single slice. In this case, if Readdir succeeds (reads all
// the way to the end of the directory), it returns the slice and a
// nil error. If it encounters an error before the end of the
// directory, Readdir returns the FileInfo read until that point
// and a non-nil error.
//
// Most clients are better served by the more efficient ReadDir method.
<原文结束>

# <翻译开始>
// Readdir 读取与文件关联的目录内容，并返回一个最多包含 n 个 FileInfo 值的切片，这些值如同通过 Lstat 返回的一样，按目录顺序排列。对同一文件的后续调用将返回更多的 FileInfo。
// 
// 若 n > 0，Readdir 最多返回 n 个 FileInfo 结构。在这种情况下，若 Readdir 返回一个空切片，它会同时返回一个非 nil 错误以解释原因。在目录末尾，错误为 io.EOF。
// 
// 若 n <= 0，Readdir 一次性返回目录中的所有 FileInfo，置于单一切片中。若 Readdir 成功（读取至目录末尾），则返回该切片及 nil 错误。若在到达目录末尾前遇到错误，Readdir 返回直至出错点已读取的 FileInfo 切片及非 nil 错误。
// 
// 对于大多数客户端而言，使用更高效的 ReadDir 方法更为适宜。
# <翻译结束>


<原文开始>
// Readdirnames reads the contents of the directory associated with file
// and returns a slice of up to n names of files in the directory,
// in directory order. Subsequent calls on the same file will yield
// further names.
//
// If n > 0, Readdirnames returns at most n names. In this case, if
// Readdirnames returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdirnames returns all the names from the directory in
// a single slice. In this case, if Readdirnames succeeds (reads all
// the way to the end of the directory), it returns the slice and a
// nil error. If it encounters an error before the end of the
// directory, Readdirnames returns the names read until that point and
// a non-nil error.
<原文结束>

# <翻译开始>
// Readdirnames读取与file关联的目录内容，并按目录顺序返回该目录中最多n个文件名的切片。
// 对同一文件进行后续调用将返回更多的名称。
// 
// 如果n > 0，Readdirnames最多返回n个名称。在这种情况下，如果Readdirnames返回一个空切片，
// 它会返回一个非nil错误来解释原因。在目录末尾，错误为io.EOF。
// 
// 如果n <= 0，Readdirnames在一个单一切片中返回目录中的所有名称。在这种情况下，若Readdirnames
// 成功（读取到目录末尾），它将返回切片和一个nil错误。如果在到达目录末尾之前遇到错误，
// Readdirnames将返回直至该点读取到的所有名称以及一个非nil错误。
# <翻译结束>


<原文开始>
// ReadDir reads the contents of the directory associated with the file f
// and returns a slice of DirEntry values in directory order.
// Subsequent calls on the same file will yield later DirEntry records in the directory.
//
// If n > 0, ReadDir returns at most n DirEntry records.
// In this case, if ReadDir returns an empty slice, it will return an error explaining why.
// At the end of a directory, the error is io.EOF.
//
// If n <= 0, ReadDir returns all the DirEntry records remaining in the directory.
// When it succeeds, it returns a nil error (not io.EOF).
<原文结束>

# <翻译开始>
// ReadDir 读取与文件 f 关联的目录内容，并以目录顺序返回 DirEntry 值的切片。
// 对同一文件的后续调用将返回目录中后面的 DirEntry 记录。
// 
// 如果 n > 0，ReadDir 最多返回 n 个 DirEntry 记录。在这种情况下，如果 ReadDir 返回一个空切片，它会返回一个解释原因的错误。在目录末尾，错误为 io.EOF。
// 
// 如果 n <= 0，ReadDir 返回目录中剩余的所有 DirEntry 记录。成功时，它返回一个 nil 错误（而非 io.EOF）。
# <翻译结束>


<原文开始>
// ReadDir reads the named directory,
// returning all its directory entries sorted by filename.
// If an error occurs reading the directory,
// ReadDir returns the entries it was able to read before the error,
// along with the error.
<原文结束>

# <翻译开始>
// ReadDir 读取指定名称的目录，
// 并按文件名排序返回其所有目录条目。
// 若在读取目录过程中发生错误，
// ReadDir 将返回在该错误发生前所成功读取的所有条目以及该错误。
# <翻译结束>

