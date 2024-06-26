// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goroot

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// Importcfg returns an importcfg file to be passed to the
// Go compiler that contains the cached paths for the .a files for the
// standard library.
// 翻译提示:func 导入配置() (字符串, 错误) {}

// ff:
func Importcfg() (string, error) {
	var icfg bytes.Buffer

	m, err := PkgfileMap()
	if err != nil {
		return "", err
	}
	fmt.Fprintf(&icfg, "# import config")
	for importPath, export := range m {
		fmt.Fprintf(&icfg, "\npackagefile %s=%s", importPath, export)
	}
	s := icfg.String()
	return s, nil
}

// 翻译提示:var (
//	标准库包文件映射 map[string]string
//	标准库包文件错误 error
//	一次性初始化 sync.Once
//)
var (
	stdlibPkgfileMap map[string]string
	stdlibPkgfileErr error
	once             sync.Once
)

// PkgfileMap returns a map of package paths to the location on disk
// of the .a file for the package.
// The caller must not modify the map.
// 翻译提示:func 包文件映射() (map[string]string, error) {}

// ff:
func PkgfileMap() (map[string]string, error) {
	once.Do(func() {
		m := make(map[string]string)
		output, err := exec.Command("go", "list", "-export", "-e", "-f", "{{.ImportPath}} {{.Export}}", "std", "cmd").Output()
		if err != nil {
			stdlibPkgfileErr = err
		}
		for _, line := range strings.Split(string(output), "\n") {
			if line == "" {
				continue
			}
			sp := strings.SplitN(line, " ", 2)
			if len(sp) != 2 {
				stdlibPkgfileErr = fmt.Errorf("determining pkgfile map: invalid line in go list output: %q", line)
				return
			}
			importPath, export := sp[0], sp[1]
			if export != "" {
				m[importPath] = export
			}
		}
		stdlibPkgfileMap = m
	})
	return stdlibPkgfileMap, stdlibPkgfileErr
}
