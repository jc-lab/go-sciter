package sciter

import "unsafe"

const (
	FALSE SBOOL = 0
	TRUE  SBOOL = 1
)

type SCITER_APP_CMD UINT

const (
	SCITER_APP_STOP     SCITER_APP_CMD = 0 // request to quit message pump loop
	SCITER_APP_LOOP     SCITER_APP_CMD = 1 // run message pump loop until SCITER_APP_STOP or main window closure
	SCITER_APP_INIT     SCITER_APP_CMD = 2 // pass argc/argv to application: p1 - argc, p2 - CHAR** argv
	SCITER_APP_SHUTDOWN SCITER_APP_CMD = 3 // free resources of the application
	SCITER_APP_RUN      SCITER_APP_CMD = 4 // scapp mode: load JS and run message pump loop until SCITER_APP_STOP or main window closure, p1 - JS url, p2 - 0 or SciterPrimordialLoader
)

type VALUE_STRING_CVT_TYPE = UINT

type VALUE_UNIT_TYPE_STRING UINT

const (
	UT_STRING_STRING = VALUE_UNIT_TYPE_STRING(0)      // string
	UT_STRING_ERROR  = VALUE_UNIT_TYPE_STRING(1)      // is an error string
	UT_STRING_SECURE = VALUE_UNIT_TYPE_STRING(2)      // secure string ("wiped" on destroy)
	UT_STRING_SYMBOL = VALUE_UNIT_TYPE_STRING(0xffff) // symbol in tiscript sense
)

func CString(s string) *byte {
	in := []byte(s)
	out := make([]byte, len(in)+1)
	copy(out, in)
	return &out[0]
}

func GoString(bs LPCSTR) string {
	if bs == nil {
		return ""
	}

	var length int
	for ptr := bs; *ptr != 0; ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1)) {
		length++
	}

	byteSlice := unsafe.Slice(bs, length)
	return string(byteSlice)
}

func GoBytes(pointer unsafe.Pointer, len INT) []byte {
	return unsafe.Slice((*byte)(pointer), len)
}

func boolToUintptr(x bool) uintptr {
	if x {
		return 1
	}
	return 0
}
