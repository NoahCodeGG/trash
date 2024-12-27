//go:build windows
// +build windows

package trash

import (
	"syscall"
	"unsafe"
)

var (
	shell32             = syscall.NewLazyDLL("shell32.dll")
	procSHFileOperation = shell32.NewProc("SHFileOperationW")
)

const (
	FO_DELETE     = 0x0003
	FOF_ALLOWUNDO = 0x0040
)

type SHFILEOPSTRUCT struct {
	hwnd                  uintptr
	wFunc                 uint32
	pFrom                 *uint16
	pTo                   *uint16
	fFlags                uint16
	fAnyOperationsAborted bool
	hNameMappings         uintptr
	lpszProgressTitle     *uint16
}

func MoveToTrash(filePath string) error {
	utf16FilePath, err := syscall.UTF16PtrFromString(filePath + "\x00")
	if err != nil {
		return err
	}

	fileOp := SHFILEOPSTRUCT{
		wFunc:  FO_DELETE,
		pFrom:  utf16FilePath,
		fFlags: FOF_ALLOWUNDO,
	}

	ret, _, err := procSHFileOperation.Call(uintptr(unsafe.Pointer(&fileOp)))
	if ret != 0 {
		return err
	}

	return nil
}
