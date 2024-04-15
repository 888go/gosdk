// 版权所有 ? 2017 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package exec

import (
	"reflect"
	"testing"
)

func TestDedupEnv(t *testing.T) {
	t.Parallel()

	tests := []struct {
		noCase  bool
		nulOK   bool
		in      []string
		want    []string
		wantErr bool
	}{
		{
			noCase: true,
			in:     []string{"k1=v1", "k2=v2", "K1=v3"},
			want:   []string{"k2=v2", "K1=v3"},
		},
		{
			noCase: false,
			in:     []string{"k1=v1", "K1=V2", "k1=v3"},
			want:   []string{"K1=V2", "k1=v3"},
		},
		{
			in:   []string{"=a", "=b", "foo", "bar"},
			want: []string{"=b", "foo", "bar"},
		},
		{
			// #49886：保留带有前导"="符号的Windows奇特键。
			noCase: true,
			in:     []string{`=C:=C:\golang`, `=D:=D:\tmp`, `=D:=D:\`},
			want:   []string{`=C:=C:\golang`, `=D:=D:\`},
		},
		{
// #52436：（暂时）保留无效的键值对条目。
// （未来可能会在某个时刻过滤掉它们或因它们而报错。）
			in:   []string{"dodgy", "entries"},
			want: []string{"dodgy", "entries"},
		},
		{
			// 过滤掉包含 NUL 的条目。
			in:      []string{"A=a\x00b", "B=b", "C\x00C=c"},
			want:    []string{"B=b"},
			wantErr: true,
		},
		{
			// Plan 9需要保留带有NUL（空字符）的环境变量（#56544）。
			nulOK: true,
			in:    []string{"path=one\x00two"},
			want:  []string{"path=one\x00two"},
		},
	}
	for _, tt := range tests {
		got, err := dedupEnvCase(tt.noCase, tt.nulOK, tt.in)
		if !reflect.DeepEqual(got, tt.want) || (err != nil) != tt.wantErr {
			t.Errorf("Dedup(%v, %q) = %q, %v; want %q, error:%v", tt.noCase, tt.in, got, err, tt.want, tt.wantErr)
		}
	}
}
