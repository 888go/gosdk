package os

import "os"

// Hostname返回内核报告的主机名。

// ff:取主机名
// err:错误
// name:主机名
func Hostname() (name string, err error) { //md5:a77a24717b6eeb11c07845ecc476ca33
	return os.Hostname()
}
