package os

// Getpagesize 返回底层系统的内存页大小。
func Getpagesize() int { //md5:d379c113740c775f0f6aaac91ecbd1be

}

// SameFile 报告 fi1 和 fi2 是否描述了同一个文件。
// 例如，在Unix系统上，这意味着两个底层结构的设备和inode字段完全相同；
// 在其他系统上，判断可能基于路径名称。
// SameFile 只适用于本包的Stat方法返回的结果。
// 在其他情况下，它会返回false。
func SameFile(fi1, fi2 FileInfo) bool { //md5:b4e76f9d966d9daea55a383bca41dc5f

}