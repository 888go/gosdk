package os


// 翻译提示:var  (
// 拷贝FileRange函数  =  poll.复制FileRange
// splice函数                  =  poll.拼接
// 发送File函数              =  poll.发送File
// )
var (
	pollCopyFileRange = poll.CopyFileRange
	pollSplice        = poll.Splice
	pollSendFile      = poll.SendFile
)// md5:4a42c678839e6a02

