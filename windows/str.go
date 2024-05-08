// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

//---build---//go:build windows

package windows

func itoa(val int) string { // 在此处执行而非使用 fmt，以避免依赖
	if val < 0 {
		return "-" + itoa(-val)
	}
	var buf [32]byte // big enough for int64
	i := len(buf) - 1
	for val >= 10 {
		buf[i] = byte(val%10 + '0')
		i--
		val /= 10
	}
	buf[i] = byte(val + '0')
	return string(buf[i:])
}
