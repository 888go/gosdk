
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T1[P any] P
// 
// func (T1[_]) m() {}
// 
// func _[P any](x *T1[P]) {
//         // x.m exists because x is of type *T1 where T1 is a defined type
//         // (even though under(T1) is a type parameter)
//         x.m()
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T1[P any] P
// 
// func (T1[_]) m() {}
// 
// func _[P any](x *T1[P]) {
//         // x.m exists because x is of type *T1 where T1 is a defined type
//         // (even though under(T1) is a type parameter)
//         x.m()
// }
# <翻译结束>


<原文开始>
        // (&x).m doesn't exist because &x is of type *P
        // and pointers to type parameters don't have methods
<原文结束>

# <翻译开始>
        // (&x).m doesn't exist because &x is of type *P
        // and pointers to type parameters don't have methods
# <翻译结束>


<原文开始>
        // x.m doesn't exists because x is of type *T2
        // and pointers to interfaces don't have methods
<原文结束>

# <翻译开始>
        // x.m doesn't exists because x is of type *T2
        // and pointers to interfaces don't have methods
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type Foo1[t any] t
// type Bar[t any] t
// 
// func (l Foo1[t]) Foo(v Barer[t]) { v.Bar(t(l)) }
// func (b *Bar[t]) Bar(l t)        { *b = Bar[t](l) }
// 
// func _[t any](f Fooer1[t]) t {
// 	var b Bar[t]
// 	f.Foo(&b)
// 	return t(b)
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type Foo1[t any] t
// type Bar[t any] t
// 
// func (l Foo1[t]) Foo(v Barer[t]) { v.Bar(t(l)) }
// func (b *Bar[t]) Bar(l t)        { *b = Bar[t](l) }
// 
// func _[t any](f Fooer1[t]) t {
// 	var b Bar[t]
// 	f.Foo(&b)
// 	return t(b)
// }
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type Fooer2[t any] interface {
// 	Foo()
// }
// 
// type Foo2[t any] t
// 
// func (f *Foo2[t]) Foo() {}
// 
// func _[t any](v t) {
// 	var f = Foo2[t](v)
// 	_ = Fooer2[t](&f)
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type Fooer2[t any] interface {
// 	Foo()
// }
// 
// type Foo2[t any] t
// 
// func (f *Foo2[t]) Foo() {}
// 
// func _[t any](v t) {
// 	var f = Foo2[t](v)
// 	_ = Fooer2[t](&f)
// }
# <翻译结束>

