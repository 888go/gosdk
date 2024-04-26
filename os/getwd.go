package os

import "os"

// Getwd 函数返回当前工作目录对应的、以根路径为起点的路径名。若由于符号链接的存在，当前目录可通过多条路径访问，Getwd 可能返回其中任意一条。

// ff:
// err:
// dir:
func Getwd() (dir string, err error) { //md5:963d89712d0463d7ffcbe8a1d9873187
	return os.Getwd()
}
