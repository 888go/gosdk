// 版权归2017年Go作者所有。保留所有权利。
// 使用此源代码受BSD风格
// 可在LICENSE文件中找到的许可协议管辖。

package windows

type MemoryBasicInformation struct {
	// 一个指向页面区域基地址的指针。
	BaseAddress uintptr
// VirtualAlloc函数分配的页面范围的基地址的指针。BaseAddress成员指向的页面包含在此分配范围内。
	AllocationBase uintptr
	// 分配该区域时所采用的内存保护选项
	AllocationProtect uint32
	PartitionId       uint16
	// 从基地址开始的区域大小，其中所有页面具有相同的属性，以字节为单位。
	RegionSize uintptr
	// 该区域页面的状态。
	State uint32
	// 区域中页面的访问保护。
	Protect uint32
	// 区域中页面的类型
	Type uint32
}
