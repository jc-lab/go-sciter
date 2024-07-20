package sciter

type POINT struct {
	X, Y int32
}
type LPPOINT *POINT

type SIZE struct {
	Cx, Cy int32
}
type LPSIZE *SIZE

type WCHAR = uint16
type LPCWSTR *WCHAR
type LPCBYTE *byte
type SBOOL int32
type LPCSTR *byte
type LPWSTR *WCHAR

type UINT uint32
type LPUINT = *UINT
type INT = int32
type FLOAT_VALUE = float64
type UINT64 = uint64
type BYTE = byte
type INT64 = int64

// windows
type UINT_PTR uintptr
type INT_PTR uintptr
type WPARAM UINT_PTR
type LONG_PTR uintptr
type LPARAM LONG_PTR
type MSG byte
type IUnknown interface{}
type LRESULT uint32

type som_asset_t interface{}
