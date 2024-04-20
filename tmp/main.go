package main

import (
	"golang.org/x/sys/windows"
)

var (
	myLibrary  = windows.NewLazyDLL("MyLibrary.dll")
	myFunction = myLibrary.NewProc("MyFunction")
)

import "golang.org/x/sys/windows"

// DLLError describes reasons for DLL load failures.
type DLLError struct {
	F windows.DLLError
}

func (e *DLLError) Error() string { //md5:a0c3c5d343d274524c4b59351d5bfed2
	return e.F.Error()
}

func (e *DLLError) Unwrap() error { //md5:499308a8661bb4bb32282dfd01bf2950
	return e.F.Unwrap()
}