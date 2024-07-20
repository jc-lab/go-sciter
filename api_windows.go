package sciter

import (
	"golang.org/x/sys/windows"
	"path/filepath"
	"sync"
	"unsafe"
)

var defaultMutex sync.Mutex
var defaultApi *SciterAPI

func GetApi() *SciterAPI {
	defaultMutex.Lock()
	defer defaultMutex.Unlock()

	if defaultApi != nil {
		return defaultApi
	}

	// default
	setDLL("")

	return defaultApi
}

func SetApi(api *SciterAPI) {
	defaultMutex.Lock()
	defer defaultMutex.Unlock()
	setApi(api)
}

func SetDLL(dir string) {
	defaultMutex.Lock()
	defer defaultMutex.Unlock()
	setDLL(dir)
}

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

func setApi(api *SciterAPI) {
	api.SciterExec(SCITER_APP_INIT, 0, 0)
	defaultApi = api
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
