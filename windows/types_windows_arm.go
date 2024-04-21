// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

package windows

type WSAData struct {
	Version      uint16
	HighVersion  uint16
	Description  [WSADESCRIPTION_LEN + 1]byte
	SystemStatus [WSASYS_STATUS_LEN + 1]byte
	MaxSockets   uint16
	MaxUdpDg     uint16
	VendorInfo   *byte
}

type Servent struct {
	Name    *byte
	Aliases **byte
	Port    uint16
	Proto   *byte
}

type JOBOBJECT_BASIC_LIMIT_INFORMATION struct {
	PerProcessUserTimeLimit int64
	PerJobUserTimeLimit     int64
	LimitFlags              uint32
	MinimumWorkingSetSize   uintptr
	MaximumWorkingSetSize   uintptr
	ActiveProcessLimit      uint32
	Affinity                uintptr
	PriorityClass           uint32
	SchedulingClass         uint32
	_                       uint32 // pad to 8 byte boundary
}
