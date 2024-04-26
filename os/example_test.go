// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package os_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/888go/gosdk/os"
	"io/fs"
	"log"
	os2 "os"
	"path/filepath"
	"sync"
	"time"
)

func ExampleOpenFile() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleOpenFile_append() {
	// 如果文件不存在，创建它，或者追加到文件末尾
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close() // 忽略错误；优先处理Write错误
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleChmod() {
	if err := os.Chmod("some-filename", 0644); err != nil {
		log.Fatal(err)
	}
}

func ExampleChtimes() {
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("some-filename", atime, mtime); err != nil {
		log.Fatal(err)
	}
}

func ExampleFileMode() {
	fi, err := os.Lstat("some-filename")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0777, etc.
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&fs.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}

func ExampleErrNotExist() {
	filename := "a-nonexistent-file"
	if _, err := os.Stat(filename); errors.Is(err, fs.ErrNotExist) {
		fmt.Println("file does not exist")
	}
	// Output:
	// file does not exist
}

func ExampleExpand() {
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))

	// Output:
	// Good morning, Gopher!
}

func ExampleExpandEnv() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	// Output:
	// gopher lives in /usr/gopher.
}

func ExampleLookupEnv() {
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

	// Output:
	// SOME_KEY=value
	// EMPTY_KEY=
	// MISSING_KEY not set
}

func ExampleGetenv() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

	// Output:
	// gopher lives in /usr/gopher.
}

func ExampleUnsetenv() {
	os.Setenv("TMPDIR", "/my/tmp")
	defer os.Unsetenv("TMPDIR")
}

func ExampleReadDir() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ExampleMkdirTemp() {
	dir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	file := filepath.Join(dir, "tmpfile")
	if err := os.WriteFile(file, []byte("content"), 0666); err != nil {
		log.Fatal(err)
	}
}

func ExampleMkdirTemp_suffix() {
	logsDir, err := os.MkdirTemp("", "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // clean up

	// 如果需要，可以通过搜索以 *-logs 结尾的目录来提前清理日志。
	globPattern := filepath.Join(os.TempDir(), "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}

func ExampleCreateTemp() {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleCreateTemp_suffix() {
	f, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleReadFile() {
	data, err := os.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)

	// Output:
	// Hello, Gophers!
}

func ExampleWriteFile() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleMkdir() {
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleMkdirAll() {
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReadlink() {
	// 首先，我们创建一个指向文件的相对符号链接。
	d, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(d)
	targetPath := filepath.Join(d, "hello.txt")
	if err := os.WriteFile(targetPath, []byte("Hello, Gophers!"), 0644); err != nil {
		log.Fatal(err)
	}
	linkPath := filepath.Join(d, "hello.link")
	if err := os.Symlink("hello.txt", filepath.Join(d, "hello.link")); err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			// 允许示例在不支持符号链接的平台上运行。
			fmt.Printf("%s links to %s\n", filepath.Base(linkPath), "hello.txt")
			return
		}
		log.Fatal(err)
	}

	// Readlink 返回传递给 os.Symlink 的相对路径。
	dst, err := os.Readlink(linkPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s links to %s\n", filepath.Base(linkPath), dst)

	var dstAbs string
	if filepath.IsAbs(dst) {
		dstAbs = dst
	} else {
		// 符号链接的目标是相对于包含该链接的目录的。
		dstAbs = filepath.Join(filepath.Dir(linkPath), dst)
	}

	// 通过比较目标与原始目标路径上的os.Stat来检查目标是否正确。
	dstInfo, err := os.Stat(dstAbs)
	if err != nil {
		log.Fatal(err)
	}
	targetInfo, err := os.Stat(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	if !os.SameFile(dstInfo, targetInfo) {
		log.Fatalf("link destination (%s) is not the same file as %s", dstAbs, targetPath)
	}

	// Output:
	// hello.link links to hello.txt
}

func ExampleUserCacheDir() {
	dir, dirErr := os.UserCacheDir()
	if dirErr == nil {
		dir = filepath.Join(dir, "ExampleUserCacheDir")
	}

	getCache := func(name string) ([]byte, error) {
		if dirErr != nil {
			return nil, &os.PathError{Op: "getCache", Path: name, Err: os2.ErrNotExist}
		}
		return os.ReadFile(filepath.Join(dir, name))
	}

	var mkdirOnce sync.Once
	putCache := func(name string, b []byte) error {
		if dirErr != nil {
			return &os.PathError{Op: "putCache", Path: name, Err: dirErr}
		}
		mkdirOnce.Do(func() {
			if err := os.MkdirAll(dir, 0700); err != nil {
				log.Printf("can't create user cache dir: %v", err)
			}
		})
		return os.WriteFile(filepath.Join(dir, name), b, 0600)
	}

	// Read and store cached data.
	// …
	_ = getCache
	_ = putCache

	// Output:
}

func ExampleUserConfigDir() {
	dir, dirErr := os.UserConfigDir()

	var (
		configPath string
		origConfig []byte
	)
	if dirErr == nil {
		configPath = filepath.Join(dir, "ExampleUserConfigDir", "example.conf")
		var err error
		origConfig, err = os.ReadFile(configPath)
		if err != nil && !os.IsNotExist(err) {
			// 用户有一个配置文件，但我们无法读取它。 // 相反于忽略他们的配置，我们应该报告这个错误。
			log.Fatal(err)
		}
	}

	// 使用并可能对配置进行更改
	config := bytes.Clone(origConfig)
	// …

	// Save changes.
	if !bytes.Equal(config, origConfig) {
		if configPath == "" {
			log.Printf("not saving config changes: %v", dirErr)
		} else {
			err := os.MkdirAll(filepath.Dir(configPath), 0700)
			if err == nil {
				err = os.WriteFile(configPath, config, 0600)
			}
			if err != nil {
				log.Printf("error saving config changes: %v", err)
			}
		}
	}

	// Output:
}
