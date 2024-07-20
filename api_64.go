//go:build amd64 || arm64
// +build amd64 arm64

package sciter

import (
	"syscall"
	"unsafe"
)

func (a *SciterAPI) SciterFindElement(hwnd HWINDOW, pt Point, phe *HELEMENT) SCDOM_RESULT {
	unsafePt := (*uintptr)(unsafe.Pointer(&pt))
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetFocusElement,
		uintptr(hwnd),
		*unsafePt,
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}
