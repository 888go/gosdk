
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2018 Go语言作者。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可证管辖。
// md5:b62d507ecbeb2e98
# <翻译结束>


<原文开始>
// Address has a City, State and a Country.
<原文结束>

# <翻译开始>
// Address 包含城市、州和国家信息。. md5:5d6891003d64916a
# <翻译结束>


<原文开始>
// Person has a Name, Age and Address.
<原文结束>

# <翻译开始>
// Person有一个Name、Age和Address属性。. md5:d29468e58cc31e7c
# <翻译结束>


<原文开始>
// GoString makes Person satisfy the GoStringer interface.
// The return value is valid Go code that can be used to reproduce the Person struct.
<原文结束>

# <翻译开始>
// GoString让Person实现了GoStringer接口。
// 返回的值是有效的Go代码，可以用来复制Person结构体。
// md5:240e6a6ef5248fc4
# <翻译结束>


<原文开始>
	// If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p1)` would be similar to
	// Person{Name:"Warren", Age:0x1f, Addr:(*main.Address)(0x10448240)}
<原文结束>

# <翻译开始>
// 如果没有实现GoString()方法，`fmt.Printf("%#v", p1)` 的输出将会类似于
// Person{Name:"Warren", Age:31, Addr:(*main.Address)(0x10448240)}
// md5:06accd0b6f883adc
# <翻译结束>


<原文开始>
	// If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p2)` would be similar to
	// Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
<原文结束>

# <翻译开始>
// 如果没有实现GoString()方法，`fmt.Printf("%#v", p2)`的输出将会类似于
// Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
// md5:df5825dd4a82752c
# <翻译结束>

