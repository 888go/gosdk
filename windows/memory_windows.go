// 版权所有 2017 The Go 作者。保留所有权利。
// 使用本源代码须遵循 BSD 风格
// 许可协议，该协议可在 LICENSE 文件中找到。

package windows

// 翻译提示:const  (
// 	内存提交                =  0x00001000
// 	内存预留              =  0x00002000
// 	内存解除提交      =  0x00004000
// 	内存释放              =  0x00008000
// 	内存重置              =  0x00080000
// 	内存自顶向下      =  0x00100000
// 	内存写入监视      =  0x00200000
// 	物理内存              =  0x00400000
// 	内存重置撤销      =  0x01000000
// 	大页内存              =  0x20000000
// 
// 	无访问权限                =  0x00000001
// 	只读权限                  =  0x00000002
// 	读写权限                  =  0x00000004
// 	写入复制权限          =  0x00000008
// 	执行权限                  =  0x00000010
// 	执行读取权限          =  0x00000020
// 	执行读写权限          =  0x00000040
// 	执行写复制权限      =  0x00000080
// 	保护页标志              =  0x00000100
// 	无缓存权限              =  0x00000200
// 	写合并权限              =  0x00000400
// 	目标页无效              =  0x40000000
// 	目标页不更新          =  0x40000000
// 
// 	配额限制硬最小禁用  =  0x00000002
// 	配额限制硬最小启用  =  0x00000001
// 	配额限制硬最大禁用  =  0x00000008
// 	配额限制硬最大启用  =  0x00000004
// )
const (
	MEM_COMMIT      = 0x00001000
	MEM_RESERVE     = 0x00002000
	MEM_DECOMMIT    = 0x00004000
	MEM_RELEASE     = 0x00008000
	MEM_RESET       = 0x00080000
	MEM_TOP_DOWN    = 0x00100000
	MEM_WRITE_WATCH = 0x00200000
	MEM_PHYSICAL    = 0x00400000
	MEM_RESET_UNDO  = 0x01000000
	MEM_LARGE_PAGES = 0x20000000

	PAGE_NOACCESS          = 0x00000001
	PAGE_READONLY          = 0x00000002
	PAGE_READWRITE         = 0x00000004
	PAGE_WRITECOPY         = 0x00000008
	PAGE_EXECUTE           = 0x00000010
	PAGE_EXECUTE_READ      = 0x00000020
	PAGE_EXECUTE_READWRITE = 0x00000040
	PAGE_EXECUTE_WRITECOPY = 0x00000080
	PAGE_GUARD             = 0x00000100
	PAGE_NOCACHE           = 0x00000200
	PAGE_WRITECOMBINE      = 0x00000400
	PAGE_TARGETS_INVALID   = 0x40000000
	PAGE_TARGETS_NO_UPDATE = 0x40000000

	QUOTA_LIMITS_HARDWS_MIN_DISABLE = 0x00000002
	QUOTA_LIMITS_HARDWS_MIN_ENABLE  = 0x00000001
	QUOTA_LIMITS_HARDWS_MAX_DISABLE = 0x00000008
	QUOTA_LIMITS_HARDWS_MAX_ENABLE  = 0x00000004
)

type MemoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	PartitionId       uint16
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}
