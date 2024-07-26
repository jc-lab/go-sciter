package sciter

import "sync"

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

func setApi(api *SciterAPI) {
	api.SciterExec(SCITER_APP_INIT, 0, 0)
	defaultApi = api
}
