package time

import "time"

const (
	seekStart   = 0
	seekCurrent = 1
	seekEnd     = 2
) // md5:489a8240bba6ec1a

// LoadLocationFromTZData 根据给定的名称，从IANA时区数据库格式的数据中初始化一个时区。
// 数据应符合标准的IANA时区文件格式（例如，Unix系统上的/etc/localtime内容）。
// md5:ece7db04102fddc3
// ff:从TZ数据加载时区
// name:
// data:
func LoadLocationFromTZData(name string, data []byte) (*Location, error) { //md5:31b2d35bb24a50e9
	l, err := time.LoadLocationFromTZData(name, data)
	if err != nil {
		return nil, err
	}
	return &Location{*l}, err
}
