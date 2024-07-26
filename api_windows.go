package sciter

import (
	"golang.org/x/sys/windows"
	"path/filepath"
	"unsafe"
)

func setDLL(dir string) {
	var dll string
	if dir == "" {
		dll = "sciter.dll"
	} else {
		filepath.Join(dir, "sciter.dll")
	}
	api, err := loadApi(dll)
	if err != nil {
		panic(err)
	}
	setApi(api)
}

func loadApi(dll string) (*SciterAPI, error) {
	mod := windows.NewLazyDLL(dll)
	err := mod.Load()
	if err != nil {
		return nil, err
	}
	ret, _, err := mod.NewProc("SciterAPI").Call()
	if ret == 0 && err != nil {
		return nil, err
	}
	api := (*SciterAPI)(unsafe.Pointer(ret))
	return api, nil
}
