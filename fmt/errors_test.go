//版权所有2018 Go语言作者。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可证管辖。
// md5:b62d507ecbeb2e98

package fmt_test

import (
	"errors"
	"github.com/888go/gosdk/fmt"
	"reflect"
	"testing"
)

func TestErrorf(t *testing.T) {
// noVetErrorf 是一个别名，用于fmt.Errorf，它不会触发vet警告对于%w格式字符串。
// md5:dd2f5d23ee1c6b2d
	noVetErrorf := fmt.Errorf

	wrapped := errors.New("inner error")
	for _, test := range []struct {
		err        error
		wantText   string
		wantUnwrap error
		wantSplit  []error
	}{{
		err:        fmt.Errorf("%w", wrapped),
		wantText:   "inner error",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("added context: %w", wrapped),
		wantText:   "added context: inner error",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%w with added context", wrapped),
		wantText:   "inner error with added context",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%s %w %v", "prefix", wrapped, "suffix"),
		wantText:   "prefix inner error suffix",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%[2]s: %[1]w", wrapped, "positional verb"),
		wantText:   "positional verb: inner error",
		wantUnwrap: wrapped,
	}, {
		err:      fmt.Errorf("%v", wrapped),
		wantText: "inner error",
	}, {
		err:      fmt.Errorf("added context: %v", wrapped),
		wantText: "added context: inner error",
	}, {
		err:      fmt.Errorf("%v with added context", wrapped),
		wantText: "inner error with added context",
	}, {
		err:      noVetErrorf("%w is not an error", "not-an-error"),
		wantText: "%!w(string=not-an-error) is not an error",
	}, {
		err:       noVetErrorf("wrapped two errors: %w %w", errString("1"), errString("2")),
		wantText:  "wrapped two errors: 1 2",
		wantSplit: []error{errString("1"), errString("2")},
	}, {
		err:       noVetErrorf("wrapped three errors: %w %w %w", errString("1"), errString("2"), errString("3")),
		wantText:  "wrapped three errors: 1 2 3",
		wantSplit: []error{errString("1"), errString("2"), errString("3")},
	}, {
		err:       noVetErrorf("wrapped nil error: %w %w %w", errString("1"), nil, errString("2")),
		wantText:  "wrapped nil error: 1 %!w(<nil>) 2",
		wantSplit: []error{errString("1"), errString("2")},
	}, {
		err:       noVetErrorf("wrapped one non-error: %w %w %w", errString("1"), "not-an-error", errString("3")),
		wantText:  "wrapped one non-error: 1 %!w(string=not-an-error) 3",
		wantSplit: []error{errString("1"), errString("3")},
	}, {
		err:       fmt.Errorf("wrapped errors out of order: %[3]w %[2]w %[1]w", errString("1"), errString("2"), errString("3")),
		wantText:  "wrapped errors out of order: 3 2 1",
		wantSplit: []error{errString("1"), errString("2"), errString("3")},
	}, {
		err:       fmt.Errorf("wrapped several times: %[1]w %[1]w %[2]w %[1]w", errString("1"), errString("2")),
		wantText:  "wrapped several times: 1 1 2 1",
		wantSplit: []error{errString("1"), errString("2")},
	}, {
		err:        fmt.Errorf("%w", nil),
		wantText:   "%!w(<nil>)",
		wantUnwrap: nil, // still nil
	}} {
		if got, want := errors.Unwrap(test.err), test.wantUnwrap; got != want {
			t.Errorf("Formatted error: %v\nerrors.Unwrap() = %v, want %v", test.err, got, want)
		}
		if got, want := splitErr(test.err), test.wantSplit; !reflect.DeepEqual(got, want) {
			t.Errorf("Formatted error: %v\nUnwrap() []error = %v, want %v", test.err, got, want)
		}
		if got, want := test.err.Error(), test.wantText; got != want {
			t.Errorf("err.Error() = %q, want %q", got, want)
		}
	}
}

func splitErr(err error) []error {
	if e, ok := err.(interface{ Unwrap() []error }); ok {
		return e.Unwrap()
	}
	return nil
}

type errString string

func (e errString) Error() string { return string(e) }
