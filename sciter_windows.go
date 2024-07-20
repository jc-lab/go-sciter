package sciter

import (
	"github.com/lxn/win"
	"unsafe"
)

// LRESULT  SciterProc (HWINDOW hwnd, UINT msg, WPARAM wParam, LPARAM lParam) ;//{ return SAPI()->SciterProc (hwnd,msg,wParam,lParam); }
// LRESULT  SciterProcND (HWINDOW hwnd, UINT msg, WPARAM wParam, LPARAM lParam, BOOL* pbHandled) ;//{ return SAPI()->SciterProcND (hwnd,msg,wParam,lParam,pbHandled); }

func ProcND(hwnd win.HWND, msg uint, wParam uintptr, lParam uintptr) (ret int, handled bool) {
	var bHandled SBOOL
	api := GetApi()
	ret = int(api.SciterProcND(HWINDOW(unsafe.Pointer(hwnd)), UINT(msg), WPARAM(wParam), LPARAM(lParam), &bHandled))
	if bHandled == 0 {
		handled = false
	} else {
		handled = true
	}
	return
}
