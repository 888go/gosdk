// 版权所有 2023 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package platform_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/888go/gosdk/exec/internal/diff"
	"github.com/888go/gosdk/exec/internal/testenv"

	"os"
	"os/exec"
	"testing"
	"text/template"
)

var flagFix = flag.Bool("fix", false, "if true, fix out-of-date generated files")

// TestGenerated 验证 zosarch.go 文件是否为最新状态，
// 若设置了 -fix 标志，则对其进行重新生成。
func TestGenerated(t *testing.T) {
	testenv.MustHaveGoRun(t)

// 在这里，我们使用 'go run cmd/dist' 而不是 'go tool dist'，以防已安装的 cmd/dist 过时或缺失。我们不希望因二进制文件过期而导致数据出现偏差。
	cmd := testenv.Command(t, "go", "run", "cmd/dist", "list", "-json", "-broken")

	// cmd/dist 需要在环境变量中明确设置 GOROOT。
	cmd.Env = append(cmd.Environ(), "GOROOT="+testenv.GOROOT(t))

	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok && len(ee.Stderr) > 0 {
			t.Logf("stderr:\n%s", ee.Stderr)
		}
		t.Fatalf("%v: %v", cmd, err)
	}

	type listEntry struct {
		GOOS, GOARCH string
		CgoSupported bool
		FirstClass   bool
		Broken       bool
	}
	var entries []listEntry
	if err := json.Unmarshal(out, &entries); err != nil {
		t.Fatal(err)
	}

	tmplOut := new(bytes.Buffer)
	tmpl := template.Must(template.New("zosarch").Parse(zosarchTmpl))
	err = tmpl.Execute(tmplOut, entries)
	if err != nil {
		t.Fatal(err)
	}

	cmd = testenv.Command(t, "gofmt")
	cmd.Stdin = bytes.NewReader(tmplOut.Bytes())
	want, err := cmd.Output()
	if err != nil {
		t.Logf("stdin:\n%s", tmplOut.Bytes())
		if ee, ok := err.(*exec.ExitError); ok && len(ee.Stderr) > 0 {
			t.Logf("stderr:\n%s", ee.Stderr)
		}
		t.Fatalf("%v: %v", cmd, err)
	}

	got, err := os.ReadFile("zosarch.go")
	if err == nil && bytes.Equal(got, want) {
		return
	}

	if !*flagFix {
		if err != nil {
			t.Log(err)
		} else {
			t.Logf("diff:\n%s", diff.Diff("zosarch.go", got, "want", want))
		}
		t.Fatalf("zosarch.go is missing or out of date; to regenerate, run\ngo generate internal/platform")
	}

	if err := os.WriteFile("zosarch.go", want, 0666); err != nil {
		t.Fatal(err)
	}
}

const zosarchTmpl = `// 代码由 go test internal/platform -fix 生成。请勿编辑。

// 若要更改此文件中的信息，请编辑 cmd/dist/build.go 中的 cgoEnabled 和/或 firstClass 映射，然后运行 'go generate internal/platform'。

package platform

// List 是所有有效 GOOS/GOARCH 组合的列表，
// 其中包括已知存在问题的端口。
var List = []OSArch{
{{range .}}	{ {{ printf "%q" .GOOS }}, {{ printf "%q" .GOARCH }} },
{{end}}
}

var distInfo = map[OSArch]osArchInfo {
{{range .}}	{ {{ printf "%q" .GOOS }}, {{ printf "%q" .GOARCH }} }:
{ {{if .CgoSupported}}CgoSupported: true, {{end}}{{if .FirstClass}}FirstClass: true, {{end}}{{if .Broken}} Broken: true, {{end}} },
{{end}}
}
`
