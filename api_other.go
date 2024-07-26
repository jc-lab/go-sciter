//go:build !windows
// +build !windows

package sciter

import (
	"fmt"
	"github.com/ebitengine/purego"
	"path/filepath"
	"unsafe"
)

func setDLL(dir string) {
	var dll string
	if dir == "" {
		dll = "libsciter.so"
	} else {
		filepath.Join(dir, "libsciter.so")
	}
	api, err := loadApi(dll)
	if err != nil {
		panic(err)
	}
	setApi(api)
}

func loadApi(dll string) (*SciterAPI, error) {
	mod, err := purego.Dlopen(dll, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return nil, err
	}
	symSciterAPI, err := purego.Dlsym(mod, "SciterAPI")
	if err != nil {
		return nil, err
	}
	ret, _, errn := purego.SyscallN(symSciterAPI)
	if ret == 0 && errn != 0 {
		return nil, fmt.Errorf("syscall failed: %v", errn)
	}
	api := (*SciterAPI)(unsafe.Pointer(ret))
	return api, nil
}
