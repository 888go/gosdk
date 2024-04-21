package os

// DirEntry 是从目录中读取的一个条目（通过使用 ReadDir 函数或 File 的 ReadDir 方法）。
type DirEntry = fs.DirEntry //md5:34ff733f235d44c78ee530088f976138



// Readdir 读取与文件关联的目录内容，并返回一个最多包含 n 个 FileInfo 值的切片，这些值如同通过 Lstat 返回的一样，按目录顺序排列。对同一文件的后续调用将返回更多的 FileInfo。
// 
// 若 n > 0，Readdir 最多返回 n 个 FileInfo 结构。在这种情况下，若 Readdir 返回一个空切片，它会同时返回一个非 nil 错误以解释原因。在目录末尾，错误为 io.EOF。
// 
// 若 n <= 0，Readdir 一次性返回目录中的所有 FileInfo，置于单一切片中。若 Readdir 成功（读取至目录末尾），则返回该切片及 nil 错误。若在到达目录末尾前遇到错误，Readdir 返回直至出错点已读取的 FileInfo 切片及非 nil 错误。
// 
// 对于大多数客户端而言，使用更高效的 ReadDir 方法更为适宜。
func (f *File) Readdir(n int) ([]FileInfo, error) { //md5:2c75130d854b4c1fda9fe4877d86013f

}

// Readdirnames读取与file关联的目录内容，并按目录顺序返回该目录中最多n个文件名的切片。
// 对同一文件进行后续调用将返回更多的名称。
// 
// 如果n > 0，Readdirnames最多返回n个名称。在这种情况下，如果Readdirnames返回一个空切片，
// 它会返回一个非nil错误来解释原因。在目录末尾，错误为io.EOF。
// 
// 如果n <= 0，Readdirnames在一个单一切片中返回目录中的所有名称。在这种情况下，若Readdirnames
// 成功（读取到目录末尾），它将返回切片和一个nil错误。如果在到达目录末尾之前遇到错误，
// Readdirnames将返回直至该点读取到的所有名称以及一个非nil错误。
func (f *File) Readdirnames(n int) (names []string, err error) { //md5:9be2f3595d8c19b3213f3c7e9125f5d9

}

// ReadDir 读取与文件 f 关联的目录内容，并以目录顺序返回 DirEntry 值的切片。
// 对同一文件的后续调用将返回目录中后面的 DirEntry 记录。
// 
// 如果 n > 0，ReadDir 最多返回 n 个 DirEntry 记录。在这种情况下，如果 ReadDir 返回一个空切片，它会返回一个解释原因的错误。在目录末尾，错误为 io.EOF。
// 
// 如果 n <= 0，ReadDir 返回目录中剩余的所有 DirEntry 记录。成功时，它返回一个 nil 错误（而非 io.EOF）。
func (f *File) ReadDir(n int) ([]DirEntry, error) { //md5:4b4bdd44564fdade96f10997d095226d

}

// ReadDir 读取指定名称的目录，
// 并按文件名排序返回其所有目录条目。
// 若在读取目录过程中发生错误，
// ReadDir 将返回在该错误发生前所成功读取的所有条目以及该错误。
func ReadDir(name string) ([]DirEntry, error) { //md5:18e885cdf820509329a5ce6895d4d4b8

}