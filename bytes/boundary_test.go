// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//go:build linux

package bytes_test

import (
	. "github.com/888go/gosdk/bytes"
	"syscall"
	"testing"
)

// 本文件测试了字节操作检查靠近页面边界数据的情况。
// 我们要确保这些操作不会读取跨越边界并引发不应有的页面错误。
// md5:75426777058c008d

// 这些测试仅在Linux上运行。被测试的代码不依赖特定操作系统，因此不需要在所有操作系统上进行测试。
// md5:88ba7e84796d0671

// dangerousSlice 返回一个切片，它立即由一个出错的页面前后跟随。
// md5:81173505fbc4ca2e
func dangerousSlice(t *testing.T) []byte {
	pagesize := syscall.Getpagesize()
	b, err := syscall.Mmap(0, 0, 3*pagesize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("mmap failed %s", err)
	}
	err = syscall.Mprotect(b[:pagesize], syscall.PROT_NONE)
	if err != nil {
		t.Fatalf("mprotect low failed %s\n", err)
	}
	err = syscall.Mprotect(b[2*pagesize:], syscall.PROT_NONE)
	if err != nil {
		t.Fatalf("mprotect high failed %s\n", err)
	}
	return b[pagesize : 2*pagesize]
}

func TestEqualNearPageBoundary(t *testing.T) {
	t.Parallel()
	b := dangerousSlice(t)
	for i := range b {
		b[i] = 'A'
	}
	for i := 0; i <= len(b); i++ {
		Equal(b[:i], b[len(b)-i:])
		Equal(b[len(b)-i:], b[:i])
	}
}

func TestIndexByteNearPageBoundary(t *testing.T) {
	t.Parallel()
	b := dangerousSlice(t)
	for i := range b {
		idx := IndexByte(b[i:], 1)
		if idx != -1 {
			t.Fatalf("IndexByte(b[%d:])=%d, want -1\n", i, idx)
		}
	}
}

func TestIndexNearPageBoundary(t *testing.T) {
	t.Parallel()
	q := dangerousSlice(t)
	if len(q) > 64 {
		// 只有在接近页面结束时才关心这个。. md5:b8599e6bc302cde8
		q = q[len(q)-64:]
	}
	b := dangerousSlice(t)
	if len(b) > 256 {
		// 只有在接近页面结束时才关心这个。. md5:b8599e6bc302cde8
		b = b[len(b)-256:]
	}
	for j := 1; j < len(q); j++ {
		q[j-1] = 1 // 差异只在最后一个字节上发现. md5:15407f2f4c5ab9e5
		for i := range b {
			idx := Index(b[i:], q[:j])
			if idx != -1 {
				t.Fatalf("Index(b[%d:], q[:%d])=%d, want -1\n", i, j, idx)
			}
		}
		q[j-1] = 0
	}

	// 测试在页边界结束时，不同对齐方式和大小的q。. md5:257221af700ea72a
	q[len(q)-1] = 1 // 差异只在最后一个字节上发现. md5:15407f2f4c5ab9e5
	for j := 0; j < len(q); j++ {
		for i := range b {
			idx := Index(b[i:], q[j:])
			if idx != -1 {
				t.Fatalf("Index(b[%d:], q[%d:])=%d, want -1\n", i, j, idx)
			}
		}
	}
	q[len(q)-1] = 0
}

func TestCountNearPageBoundary(t *testing.T) {
	t.Parallel()
	b := dangerousSlice(t)
	for i := range b {
		c := Count(b[i:], []byte{1})
		if c != 0 {
			t.Fatalf("Count(b[%d:], {1})=%d, want 0\n", i, c)
		}
		c = Count(b[:i], []byte{0})
		if c != i {
			t.Fatalf("Count(b[:%d], {0})=%d, want %d\n", i, c, i)
		}
	}
}
