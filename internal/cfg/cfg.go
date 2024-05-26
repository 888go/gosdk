// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cfg holds configuration shared by the Go command and internal/testenv.
// Definitions that don't need to be exposed outside of cmd/go should be in
// cmd/go/internal/cfg instead of this package.
package cfg

// KnownEnv is a list of environment variables that affect the operation
// of the Go command.
// 翻译提示:const 知道的环境变量 = `
//
//// 下面的常量假设为配置文件类型的枚举
//const (
//    JSON配置 = "json"
//    YAML配置 = "yaml"
//    TOML配置 = "toml"
//    PROTOBUF配置 = "protobuf"
//)
//
//// 以下为配置加载位置的常量
//const (
//    默认配置路径 = "default_config_path"
//    用户配置路径 = "user_config_path"
//    系统配置路径 = "system_config_path"
//    环境变量配置 = "env_config"
//    命令行参数配置 = "cmd_args_config"
//)
//
//// 错误码常量
//const (
//    加载配置失败 = "load_config_failure"
//    解析配置错误 = "parse_config_error"
//    找不到配置文件 = "config_file_not_found"
//)
const KnownEnv = `
	AR
	CC
	CGO_CFLAGS
	CGO_CFLAGS_ALLOW
	CGO_CFLAGS_DISALLOW
	CGO_CPPFLAGS
	CGO_CPPFLAGS_ALLOW
	CGO_CPPFLAGS_DISALLOW
	CGO_CXXFLAGS
	CGO_CXXFLAGS_ALLOW
	CGO_CXXFLAGS_DISALLOW
	CGO_ENABLED
	CGO_FFLAGS
	CGO_FFLAGS_ALLOW
	CGO_FFLAGS_DISALLOW
	CGO_LDFLAGS
	CGO_LDFLAGS_ALLOW
	CGO_LDFLAGS_DISALLOW
	CXX
	FC
	GCCGO
	GO111MODULE
	GO386
	GOAMD64
	GOARCH
	GOARM
	GOBIN
	GOCACHE
	GOENV
	GOEXE
	GOEXPERIMENT
	GOFLAGS
	GOGCCFLAGS
	GOHOSTARCH
	GOHOSTOS
	GOINSECURE
	GOMIPS
	GOMIPS64
	GOMODCACHE
	GONOPROXY
	GONOSUMDB
	GOOS
	GOPATH
	GOPPC64
	GOPRIVATE
	GOPROXY
	GOROOT
	GOSUMDB
	GOTMPDIR
	GOTOOLDIR
	GOVCS
	GOWASM
	GOWORK
	GO_EXTLINK_ENABLED
	PKG_CONFIG
`
