//go:build 386 || arm
// +build 386 arm

package sciter

import (
	"github.com/ebitengine/purego"
	"unsafe"
)

func (a *SciterAPI) SciterFindElement(hwnd HWINDOW, pt Point, phe *HELEMENT) SCDOM_RESULT {
	unsafePt := (*uintptr)(unsafe.Pointer(&pt))
	ret, _, _ := purego.SyscallN(
		a.fnSciterGetFocusElement,
		uintptr(hwnd),
		unsafePt[0],
		unsafePt[1],
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}
