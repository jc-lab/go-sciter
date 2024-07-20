package sciter

import "C"
import (
	"fmt"
	"unsafe"
)

/** Resource data type.
 *  Used by SciterDataReadyAsync() function.
 **/
type SciterResourceType uint32

// typedef enum SciterResourceType
const (
	RT_DATA_HTML SciterResourceType = iota
	RT_DATA_IMAGE
	RT_DATA_STYLE
	RT_DATA_CURSOR
	RT_DATA_SCRIPT
	RT_DATA_RAW
	RT_DATA_FONT
	RT_DATA_SOUND // wav bytes

	RT_DATA_FORCE_DWORD = 0xffffffff
)

type REQUEST_RESULT int32

// enum REQUEST_RESULT
const (
	REQUEST_PANIC = iota - 1 // e.g. not enough memory
	REQUEST_OK
	REQUEST_BAD_PARAM    // bad parameter
	REQUEST_FAILURE      // operation failed, e.g. index out of bounds
	REQUEST_NOTSUPPORTED // the platform does not support requested feature
)

// enum REQUEST_RQ_TYPE
const (
	RRT_GET = 1 + iota
	RRT_POST
	RRT_PUT
	RRT_DELETE

	RRT_FORCE_DWORD = 0xFFFFFFFF
)

// enum REQUEST_STATE
const (
	RS_PENDING = iota
	RS_SUCCESS // completed successfully
	RS_FAILURE // completed with failure

	RS_FORCE_DWORD = 0xFFFFFFFF
)

type requestError struct {
	Result  REQUEST_RESULT
	Message string
}

func (e *requestError) Error() string {
	return fmt.Sprintf("%s: %s", e.Result.String(), e.Message)
}

func newRequestError(ret REQUEST_RESULT, msg string) *requestError {
	return &requestError{
		Result:  ret,
		Message: msg,
	}
}

func wrapRequestResult(r REQUEST_RESULT, msg string) error {
	if r == REQUEST_RESULT(REQUEST_OK) {
		return nil
	}
	return newRequestError(REQUEST_RESULT(r), msg)
}

var (
	BAD_HREQUEST = HREQUEST(unsafe.Pointer(uintptr(0)))
)

// Request represents a request issued by sciter
// e.g. el.request(...) or view.request(...)
type Request struct {
	handle HREQUEST
}

//// WrapRequest wraps HREQUEST to a go side *Request, doing RequestUse/RequestUnUse automatically
//func WrapRequest(requestId HREQUEST) *Request {
//	r := &Request{
//		handle: requestId,
//	}
//	r.use()
//	runtime.SetFinalizer(r, (*Request).finalize)
//	return r
//}
//
//func (r *Request) use() error {
//	ret := RequestUse(r.handle)
//	return wrapRequestResult(ret, "RequestUse")
//}
//
//func (r *Request) unUse() error {
//	ret := RequestUnUse(r.handle)
//	return wrapRequestResult(ret, "RequestUnUse")
//}
//
//func (r *Request) finalize() {
//	r.unUse()
//	r.handle = BAD_HREQUEST
//}
//
//func (r *Request) Url() (string, error) {
//	var url string
//	// args
//	curl := LPVOID(unsafe.Pointer(&url))
//	// cgo call
//	ret := RequestUrl(r.handle, lpcstr_receiver, curl)
//	return url, wrapRequestResult(ret, "RequestUrl")
//}
//
//func (r *Request) ContentUrl() (string, error) {
//	var url string
//	// args
//	curl := LPVOID(unsafe.Pointer(&url))
//	// cgo call
//	ret := defaultApi.RequestContentUrl(r.handle, lpcstr_receiver, curl)
//	return url, wrapRequestResult(ret, "RequestContentUrl")
//}
//
//func (r *Request) RequestType() (uint, error) {
//	var requestType uint
//	// args
//	crequestType := (*REQUEST_RQ_TYPE)(unsafe.Pointer(&requestType))
//	// cgo call
//	ret := RequestGetRequestType(r.handle, crequestType)
//	return requestType, wrapRequestResult(ret, "RequestGetRequestType")
//}
//
//func (r *Request) RequestedDataType() (SciterResourceType, error) {
//	var resourceType SciterResourceType
//	// args
//	cresourceType := (*SciterResourceType)(unsafe.Pointer(&resourceType))
//	// cgo call
//	ret := RequestGetRequestedDataType(r.handle, cresourceType)
//	return resourceType, wrapRequestResult(ret, "RequestGetRequestedDataType")
//}
//
//func (r *Request) ReceivedDataType() (string, error) {
//	var dataType string
//	// args
//	cdataType := LPVOID(unsafe.Pointer(&dataType))
//	// cgo call
//	ret := RequestGetReceivedDataType(r.handle, lpcstr_receiver, cdataType)
//	return dataType, wrapRequestResult(ret, "RequestGetReceivedDataType")
//}
//
//func (r *Request) NumberOfParameters() (uint, error) {
//	var numParams uint
//	// args
//	cnumParams := (*UINT)(unsafe.Pointer(&numParams))
//	// cgo call
//	ret := RequestGetNumberOfParameters(r.handle, cnumParams)
//	return numParams, wrapRequestResult(ret, "RequestGetNumberOfParameters")
//}
//
//func (r *Request) NthParameterName(idx uint) (string, error) {
//	var name string
//	// args
//	cname := LPVOID(unsafe.Pointer(&name))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthParameterName(r.handle, cidx, lpcwstr_receiver, cname)
//	return name, wrapRequestResult(ret, "RequestGetNthParameterName")
//}
//
//func (r *Request) NthParameterValue(idx uint) (string, error) {
//	var value string
//	// args
//	cvalue := LPVOID(unsafe.Pointer(&value))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthParameterValue(r.handle, cidx, lpcwstr_receiver, cvalue)
//	return value, wrapRequestResult(ret, "RequestGetNthParameterValue")
//}
//
//func (r *Request) Times() (time.Time, time.Time, error) {
//	var started, ended uint
//	var tStarted, tEnded time.Time
//	// args
//	cstarted := (*UINT)(unsafe.Pointer(&started))
//	cended := (*UINT)(unsafe.Pointer(&ended))
//	// cgo call
//	ret := RequestGetTimes(r.handle, cstarted, cended)
//	tStarted = time.Unix(int64(started), 0)
//	tEnded = time.Unix(int64(ended), 0)
//	return tStarted, tEnded, wrapRequestResult(ret, "RequestGetTimes")
//}
//
//func (r *Request) NumberOfRqHeaders() (uint, error) {
//	var num uint
//	// args
//	cnum := (*UINT)(unsafe.Pointer(&num))
//	// cgo call
//	ret := RequestGetNumberOfRqHeaders(r.handle, cnum)
//	return num, wrapRequestResult(ret, "RequestGetNumberOfRqHeaders")
//}
//
//func (r *Request) NthRqHeaderName(idx uint) (string, error) {
//	var name string
//	// args
//	cname := LPVOID(unsafe.Pointer(&name))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthRqHeaderName(r.handle, cidx, lpcwstr_receiver, cname)
//	return name, wrapRequestResult(ret, "RequestGetNthRqHeaderName")
//}
//
//func (r *Request) NthRqHeaderValue(idx uint) (string, error) {
//	var val string
//	// args
//	cvalue := LPVOID(unsafe.Pointer(&val))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthRqHeaderValue(r.handle, cidx, lpcwstr_receiver, cvalue)
//	return val, wrapRequestResult(ret, "RequestGetNthRqHeaderValue")
//}
//
//func (r *Request) NumberOfRspHeaders() (uint, error) {
//	var num uint
//	// args
//	cnum := (*UINT)(unsafe.Pointer(&num))
//	// cgo call
//	ret := RequestGetNumberOfRspHeaders(r.handle, cnum)
//	return num, wrapRequestResult(ret, "RequestGetNumberOfRspHeaders")
//}
//
//func (r *Request) NthRspHeaderName(idx uint) (string, error) {
//	var name string
//	// args
//	cname := LPVOID(unsafe.Pointer(&name))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthRspHeaderName(r.handle, cidx, lpcwstr_receiver, cname)
//	return name, wrapRequestResult(ret, "RequestGetNthRspHeaderName")
//}
//
//func (r *Request) NthRspHeaderValue(idx uint) (string, error) {
//	var val string
//	// args
//	cvalue := LPVOID(unsafe.Pointer(&val))
//	cidx := UINT(idx)
//	// cgo call
//	ret := RequestGetNthRspHeaderValue(r.handle, cidx, lpcwstr_receiver, cvalue)
//	return val, wrapRequestResult(ret, "RequestGetNthRspHeaderValue")
//}
//
//func (r *Request) CompletionStatus() (uint, uint, error) {
//	var state, status uint
//	// args
//	cstate := (*REQUEST_STATE)(unsafe.Pointer(&state))
//	cstatus := (*UINT)(unsafe.Pointer(&status))
//	// cgo call
//	ret := RequestGetCompletionStatus(r.handle, cstate, cstatus)
//	return state, status, wrapRequestResult(ret, "RequestGetCompletionStatus")
//}
//
//func (r *Request) ProxyHost() (string, error) {
//	var host string
//	// args
//	chost := LPVOID(unsafe.Pointer(&host))
//	// cgo call
//	ret := RequestGetProxyHost(r.handle, lpcstr_receiver, chost)
//	return host, wrapRequestResult(ret, "RequestGetProxyHost")
//}
//
//func (r *Request) ProxyPort() (uint, error) {
//	var portno uint
//	// args
//	cportno := (*UINT)(unsafe.Pointer(&portno))
//	// cgo call
//	ret := RequestGetProxyPort(r.handle, cportno)
//	return portno, wrapRequestResult(ret, "RequestGetProxyPort")
//}
//
//func (r *Request) SetSucceeded(status uint, data []byte) error {
//	// args
//	cstatus := UINT(status)
//	var pData LPCBYTE
//	if len(data) > 0 {
//		pData = LPCBYTE(unsafe.Pointer(&data[0]))
//	}
//	cdataLength := UINT(len(data))
//
//	//cgo call
//	ret := RequestSetSucceeded(r.handle, cstatus, pData, cdataLength)
//	return wrapRequestResult(ret, "RequestSetSucceeded")
//}
//
//func (r *Request) SetFailed(status uint, data []byte) error {
//	// args
//	cstatus := UINT(status)
//	var pData LPCBYTE
//	if len(data) > 0 {
//		pData = LPCBYTE(unsafe.Pointer(&data[0]))
//	}
//	cdataLength := UINT(len(data))
//
//	//cgo call
//	ret := RequestSetFailed(r.handle, cstatus, pData, cdataLength)
//	return wrapRequestResult(ret, "RequestSetFailed")
//}
//
//func (r *Request) AppendDataChunk(data []byte) error {
//	// args
//	var pData LPCBYTE
//	if len(data) > 0 {
//		pData = LPCBYTE(unsafe.Pointer(&data[0]))
//	}
//	cdataLength := UINT(len(data))
//
//	//cgo call
//	ret := RequestAppendDataChunk(r.handle, pData, cdataLength)
//	return wrapRequestResult(ret, "RequestAppendDataChunk")
//}
//
//func (r *Request) SetRqHeader(name, value string) error {
//	// args
//	cname := LPCWSTR(unsafe.Pointer(StringToWcharPtr(name)))
//	cvalue := LPCWSTR(unsafe.Pointer(StringToWcharPtr(value)))
//	// cgo call
//	ret := RequestSetRqHeader(r.handle, cname, cvalue)
//	return wrapRequestResult(ret, "RequestSetRqHeader")
//}
//
//func (r *Request) SetRspHeader(name, value string) error {
//	// args
//	cname := LPCWSTR(unsafe.Pointer(StringToWcharPtr(name)))
//	cvalue := LPCWSTR(unsafe.Pointer(StringToWcharPtr(value)))
//	// cgo call
//	ret := RequestSetRspHeader(r.handle, cname, cvalue)
//	return wrapRequestResult(ret, "RequestSetRspHeader")
//}
//
//func (r *Request) SetReceivedDataType(dataType string) error {
//	// args
//	cdataType := LPCSTR(unsafe.Pointer(StringToBytePtr(dataType)))
//	// cgo call
//	ret := RequestSetReceivedDataType(r.handle, cdataType)
//	return wrapRequestResult(ret, "RequestSetReceivedDataType")
//}
//
//func (r *Request) RequestSetReceivedDataEncoding(encoding string) error {
//	// args
//	cencoding := LPCSTR(unsafe.Pointer(StringToBytePtr(encoding)))
//	// cgo call
//	ret := RequestSetReceivedDataEncoding(r.handle, cencoding)
//	return wrapRequestResult(ret, "RequestSetReceivedDataEncoding")
//}
//
//func (r *Request) Data() ([]byte, error) {
//	var buf []byte
//	// args
//	cbuf := LPVOID(unsafe.Pointer(&buf))
//	// cgo call
//	ret := RequestGetData(r.handle, lpcbyte_receiver, cbuf)
//	return buf, wrapRequestResult(ret, "RequestGetData")
//}
