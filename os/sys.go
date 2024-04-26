package os

import "os"

// Hostname返回内核报告的主机名。

// ff:
// err:
// name:
func Hostname() (name string, err error) { //md5:a77a24717b6eeb11c07845ecc476ca33
	return os.Hostname()
}
