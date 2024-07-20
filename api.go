package sciter

import (
	"syscall"
	"unsafe"
)

type LPVOID = uintptr
type HRANGE uintptr
type HELEMENT uintptr
type HNODE uintptr
type HSARCHIVE uintptr
type HPOSITION struct {
	Hn  HNODE
	Pos int32
}
type HWINDOW uintptr
type HREQUEST uintptr

type SciterAPI struct {
	Version                UINT // API_VERSION
	fnSciterClassName      uintptr
	fnSciterVersion        uintptr
	fnSciterDataReady      uintptr
	fnSciterDataReadyAsync uintptr

	// #if defined(WINDOWS) && !defined(WINDOWLESS)
	fnSciterProc   uintptr
	fnSciterProcND uintptr

	// #endif
	fnSciterLoadFile        uintptr
	fnSciterLoadHtml        uintptr
	fnSciterSetCallback     uintptr
	fnSciterSetMasterCSS    uintptr
	fnSciterAppendMasterCSS uintptr
	fnSciterSetCSS          uintptr
	fnSciterSetMediaType    uintptr
	fnSciterSetMediaVars    uintptr
	fnSciterGetMinWidth     uintptr
	fnSciterGetMinHeight    uintptr
	fnSciterCall            uintptr
	fnSciterEval            uintptr
	fnSciterUpdateWindow    uintptr

	// #if defined(WINDOWS) && !defined(WINDOWLESS)
	fnSciterTranslateMessage uintptr

	// #endif
	fnSciterSetOption      uintptr
	fnSciterGetPPI         uintptr
	fnSciterGetViewExpando uintptr

	// if defined(WINDOWS) && !defined(WINDOWLESS)
	fnSciterRenderD2D  uintptr
	fnSciterD2DFactory uintptr
	fnSciterDWFactory  uintptr

	// endif
	fnSciterGraphicsCaps uintptr
	fnSciterSetHomeURL   uintptr

	// if defined(OSX) && !defined(WINDOWLESS)
	fnSciterCreateNSView uintptr // returns NSView*

	// endif
	// if defined(LINUX) && !defined(WINDOWLESS)
	fnSciterCreateWidget uintptr // returns GtkWidget

	// endif
	// if !defined(WINDOWLESS)
	fnSciterCreateWindow uintptr

	// endif
	fnSciterSetupDebugOutput uintptr

	//|
	//| DOM Element API
	//|
	fnSciter_UseElement               uintptr
	fnSciter_UnuseElement             uintptr
	fnSciterGetRootElement            uintptr
	fnSciterGetFocusElement           uintptr
	fnSciterFindElement               uintptr
	fnSciterGetChildrenCount          uintptr
	fnSciterGetNthChild               uintptr
	fnSciterGetParentElement          uintptr
	fnSciterGetElementHtmlCB          uintptr
	fnSciterGetElementTextCB          uintptr
	fnSciterSetElementText            uintptr
	fnSciterGetAttributeCount         uintptr
	fnSciterGetNthAttributeNameCB     uintptr
	fnSciterGetNthAttributeValueCB    uintptr
	fnSciterGetAttributeByNameCB      uintptr
	fnSciterSetAttributeByName        uintptr
	fnSciterClearAttributes           uintptr
	fnSciterGetElementIndex           uintptr
	fnSciterGetElementType            uintptr
	fnSciterGetElementTypeCB          uintptr
	fnSciterGetStyleAttributeCB       uintptr
	fnSciterSetStyleAttribute         uintptr
	fnSciterGetElementLocation        uintptr
	fnSciterScrollToView              uintptr
	fnSciterUpdateElement             uintptr
	fnSciterRefreshElementArea        uintptr
	fnSciterSetCapture                uintptr
	fnSciterReleaseCapture            uintptr
	fnSciterGetElementHwnd            uintptr
	fnSciterCombineURL                uintptr
	fnSciterSelectElements            uintptr
	fnSciterSelectElementsW           uintptr
	fnSciterSelectParent              uintptr
	fnSciterSelectParentW             uintptr
	fnSciterSetElementHtml            uintptr
	fnSciterGetElementUID             uintptr
	fnSciterGetElementByUID           uintptr
	fnSciterShowPopup                 uintptr
	fnSciterShowPopupAt               uintptr
	fnSciterHidePopup                 uintptr
	fnSciterGetElementState           uintptr
	fnSciterSetElementState           uintptr
	fnSciterCreateElement             uintptr
	fnSciterCloneElement              uintptr
	fnSciterInsertElement             uintptr
	fnSciterDetachElement             uintptr
	fnSciterDeleteElement             uintptr
	fnSciterSetTimer                  uintptr
	fnSciterDetachEventHandler        uintptr
	fnSciterAttachEventHandler        uintptr
	fnSciterWindowAttachEventHandler  uintptr
	fnSciterWindowDetachEventHandler  uintptr
	fnSciterSendEvent                 uintptr
	fnSciterPostEvent                 uintptr
	fnSciterCallBehaviorMethod        uintptr
	fnSciterRequestElementData        uintptr
	fnSciterHttpRequest               uintptr
	fnSciterGetScrollInfo             uintptr
	fnSciterSetScrollPos              uintptr
	fnSciterGetElementIntrinsicWidths uintptr
	fnSciterGetElementIntrinsicHeight uintptr
	fnSciterIsElementVisible          uintptr
	fnSciterIsElementEnabled          uintptr
	fnSciterSortElements              uintptr
	fnSciterSwapElements              uintptr
	fnSciterTraverseUIEvent           uintptr
	fnSciterCallScriptingMethod       uintptr
	fnSciterCallScriptingFunction     uintptr
	fnSciterEvalElementScript         uintptr
	fnSciterAttachHwndToElement       uintptr
	fnSciterControlGetType            uintptr
	fnSciterGetValue                  uintptr
	fnSciterSetValue                  uintptr
	fnSciterGetExpando                uintptr
	fnSciterGetObject                 uintptr
	fnSciterGetElementNamespace       uintptr
	fnSciterGetHighlightedElement     uintptr
	fnSciterSetHighlightedElement     uintptr

	//|
	//| DOM Node API
	//|
	fnSciterNodeAddRef          uintptr
	fnSciterNodeRelease         uintptr
	fnSciterNodeCastFromElement uintptr
	fnSciterNodeCastToElement   uintptr
	fnSciterNodeFirstChild      uintptr
	fnSciterNodeLastChild       uintptr
	fnSciterNodeNextSibling     uintptr
	fnSciterNodePrevSibling     uintptr
	fnSciterNodeParent          uintptr
	fnSciterNodeNthChild        uintptr
	fnSciterNodeChildrenCount   uintptr
	fnSciterNodeType            uintptr
	fnSciterNodeGetText         uintptr
	fnSciterNodeSetText         uintptr
	fnSciterNodeInsert          uintptr
	fnSciterNodeRemove          uintptr
	fnSciterCreateTextNode      uintptr
	fnSciterCreateCommentNode   uintptr

	//|
	//| Value API
	//|
	fnValueInit               uintptr
	fnValueClear              uintptr
	fnValueCompare            uintptr
	fnValueCopy               uintptr
	fnValueIsolate            uintptr
	fnValueType               uintptr
	fnValueStringData         uintptr
	fnValueStringDataSet      uintptr
	fnValueIntData            uintptr
	fnValueIntDataSet         uintptr
	fnValueInt64Data          uintptr
	fnValueInt64DataSet       uintptr
	fnValueFloatData          uintptr
	fnValueFloatDataSet       uintptr
	fnValueBinaryData         uintptr
	fnValueBinaryDataSet      uintptr
	fnValueElementsCount      uintptr
	fnValueNthElementValue    uintptr
	fnValueNthElementValueSet uintptr
	fnValueNthElementKey      uintptr
	fnValueEnumElements       uintptr
	fnValueSetValueToKey      uintptr
	fnValueGetValueOfKey      uintptr
	fnValueToString           uintptr
	fnValueFromString         uintptr
	fnValueInvoke             uintptr
	fnValueNativeFunctorSet   uintptr
	fnValueIsNativeFunctor    uintptr

	// used to be script VM API
	Reserved1                LPVOID
	Reserved2                LPVOID
	Reserved3                LPVOID
	Reserved4                LPVOID
	fnSciterOpenArchive      uintptr
	fnSciterGetArchiveItem   uintptr
	fnSciterCloseArchive     uintptr
	fnSciterFireEvent        uintptr
	fnSciterGetCallbackParam uintptr
	fnSciterPostCallback     uintptr
	fnGetSciterGraphicsAPI   uintptr
	fnGetSciterRequestAPI    uintptr

	// if defined(WINDOWS) && !defined(WINDOWLESS)
	fnSciterCreateOnDirectXWindow  uintptr // IDXGISwapChain
	fnSciterRenderOnDirectXWindow  uintptr
	fnSciterRenderOnDirectXTexture uintptr // IDXGISurface

	// endif
	fnSciterProcX              uintptr // returns TRUE if handled
	fnSciterAtomValue          uintptr //
	fnSciterAtomNameCB         uintptr
	fnSciterSetGlobalAsset     uintptr
	fnSciterGetElementAsset    uintptr
	fnSciterSetVariable        uintptr
	fnSciterGetVariable        uintptr
	fnSciterElementUnwrap      uintptr
	fnSciterElementWrap        uintptr
	fnSciterNodeUnwrap         uintptr
	fnSciterNodeWrap           uintptr
	fnSciterReleaseGlobalAsset uintptr
	fnSciterExec               uintptr
	fnSciterWindowExec         uintptr
}

func (a *SciterAPI) SciterClassName() LPCWSTR {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterClassName,
	)
	return LPCWSTR(unsafe.Pointer(ret))
}

func (a *SciterAPI) SciterVersion(n UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterVersion,
		uintptr(n),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterDataReady(hwnd HWINDOW, uri LPCWSTR, data LPCBYTE, dataLength UINT) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDataReady,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(uri)),
		uintptr(unsafe.Pointer(data)),
		uintptr(dataLength),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterDataReadyAsync(hwnd HWINDOW, uri LPCWSTR, data LPCBYTE, dataLength UINT, requestId LPVOID) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDataReadyAsync,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(uri)),
		uintptr(unsafe.Pointer(data)),
		uintptr(dataLength),
		uintptr(unsafe.Pointer(requestId)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterProc(hwnd HWINDOW, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterProc,
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(unsafe.Pointer(lParam)),
	)
	return LRESULT(ret)
}

func (a *SciterAPI) SciterProcND(hwnd HWINDOW, msg UINT, wParam WPARAM, lParam LPARAM, pbHandled *SBOOL) LRESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterProcND,
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(unsafe.Pointer(lParam)),
		uintptr(unsafe.Pointer(pbHandled)),
	)
	return LRESULT(ret)
}

func (a *SciterAPI) SciterLoadFile(hWndSciter HWINDOW, filename LPCWSTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterLoadFile,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(filename)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterLoadHtml(hWndSciter HWINDOW, html LPCBYTE, htmlSize UINT, baseUrl LPCWSTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterLoadHtml,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(html)),
		uintptr(htmlSize),
		uintptr(unsafe.Pointer(baseUrl)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetCallback(hWndSciter HWINDOW, cb uintptr, cbParam LPVOID) {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetCallback,
		uintptr(hWndSciter),
		uintptr(cb),
		uintptr(unsafe.Pointer(cbParam)),
	)
	_ = ret
}

func (a *SciterAPI) SciterSetMasterCSS(utf8 LPCBYTE, numBytes UINT) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetMasterCSS,
		uintptr(unsafe.Pointer(utf8)),
		uintptr(numBytes),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterAppendMasterCSS(utf8 LPCBYTE, numBytes UINT) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterAppendMasterCSS,
		uintptr(unsafe.Pointer(utf8)),
		uintptr(numBytes),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetCSS(hWndSciter HWINDOW, utf8 LPCBYTE, numBytes UINT, baseUrl LPCWSTR, mediaType LPCWSTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetCSS,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(utf8)),
		uintptr(numBytes),
		uintptr(unsafe.Pointer(baseUrl)),
		uintptr(unsafe.Pointer(mediaType)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetMediaType(hWndSciter HWINDOW, mediaType LPCWSTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetMediaType,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(mediaType)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetMediaVars(hWndSciter HWINDOW) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetMediaVars,
		uintptr(hWndSciter),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterGetMinWidth(hWndSciter HWINDOW) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetMinWidth,
		uintptr(hWndSciter),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterGetMinHeight(hWndSciter HWINDOW, width UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetMinHeight,
		uintptr(hWndSciter),
		uintptr(width),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterCall(hWnd HWINDOW, functionName LPCSTR, argc UINT /* const */, argv *SCITER_VALUE, retval *SCITER_VALUE) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCall,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(functionName)),
		uintptr(argc),
		uintptr(unsafe.Pointer(argv)),
		uintptr(unsafe.Pointer(retval)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterEval(hwnd HWINDOW, script LPCWSTR, scriptLength UINT, pretval *SCITER_VALUE) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterEval,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(script)),
		uintptr(scriptLength),
		uintptr(unsafe.Pointer(pretval)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterUpdateWindow(hwnd HWINDOW) {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterUpdateWindow,
		uintptr(hwnd),
	)
	_ = ret
}

func (a *SciterAPI) SciterTranslateMessage(lpMsg *MSG) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterTranslateMessage,
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetOption(hWnd HWINDOW, option UINT, value UINT_PTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetOption,
		uintptr(hWnd),
		uintptr(option),
		uintptr(value),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterGetPPI(hWndSciter HWINDOW, px *UINT, py *UINT) {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetPPI,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(px)),
		uintptr(unsafe.Pointer(py)),
	)
	_ = ret
}

func (a *SciterAPI) SciterGetViewExpando(hwnd HWINDOW, pval *Value) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetViewExpando,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pval)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterRenderD2D(hWndSciter HWINDOW, prt *IUnknown /*ID2D1RenderTarget**/) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterRenderD2D,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(prt)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterD2DFactory(ppf **IUnknown /*ID2D1Factory ***/) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterD2DFactory,
		uintptr(unsafe.Pointer(ppf)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterDWFactory(ppf **IUnknown /*IDWriteFactory ***/) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDWFactory,
		uintptr(unsafe.Pointer(ppf)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterGraphicsCaps(pcaps LPUINT) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGraphicsCaps,
		uintptr(unsafe.Pointer(pcaps)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetHomeURL(hWndSciter HWINDOW, baseUrl LPCWSTR) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetHomeURL,
		uintptr(hWndSciter),
		uintptr(unsafe.Pointer(baseUrl)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterCreateNSView(frame *Rect) HWINDOW {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateNSView,
		uintptr(unsafe.Pointer(frame)),
	)
	return HWINDOW(ret)
}

func (a *SciterAPI) SciterCreateWidget(frame *Rect) HWINDOW {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateWidget,
		uintptr(unsafe.Pointer(frame)),
	)
	return HWINDOW(ret)
}

func (a *SciterAPI) SciterCreateWindow(creationFlags UINT, frame *Rect, delegate uintptr, delegateParam uintptr, parent HWINDOW) HWINDOW {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateWindow,
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(frame)),
		uintptr(delegate),
		uintptr(delegateParam),
		uintptr(parent),
	)
	return HWINDOW(ret)
}

func (a *SciterAPI) SciterSetupDebugOutput(hwndOrNull HWINDOW) {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetupDebugOutput,
		uintptr(hwndOrNull),
	)
	_ = ret
}

func (a *SciterAPI) Sciter_UseElement(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciter_UseElement,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) Sciter_UnuseElement(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciter_UnuseElement,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetRootElement(hwnd HWINDOW, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetRootElement,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetFocusElement(hwnd HWINDOW, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetFocusElement,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetChildrenCount(he HELEMENT, count *UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetChildrenCount,
		uintptr(he),
		uintptr(unsafe.Pointer(count)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetNthChild(he HELEMENT, n UINT, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetNthChild,
		uintptr(he),
		uintptr(n),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetParentElement(he HELEMENT, p_parent_he *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetParentElement,
		uintptr(he),
		uintptr(unsafe.Pointer(p_parent_he)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementHtmlCB(he HELEMENT, outer SBOOL, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementHtmlCB,
		uintptr(he),
		uintptr(outer),
		uintptr(unsafe.Pointer(rcv)),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementTextCB(he HELEMENT, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementTextCB,
		uintptr(he),
		uintptr(rcv),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetElementText(he HELEMENT, utf16 LPCWSTR, length UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetElementText,
		uintptr(he),
		uintptr(unsafe.Pointer(utf16)),
		uintptr(length),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetAttributeCount(he HELEMENT, p_count LPUINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetAttributeCount,
		uintptr(he),
		uintptr(unsafe.Pointer(p_count)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetNthAttributeNameCB(he HELEMENT, n UINT, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetNthAttributeNameCB,
		uintptr(he),
		uintptr(n),
		uintptr(unsafe.Pointer(rcv)),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetNthAttributeValueCB(he HELEMENT, n UINT, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetNthAttributeValueCB,
		uintptr(he),
		uintptr(n),
		uintptr(rcv),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetAttributeByNameCB(he HELEMENT, name LPCSTR, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetAttributeByNameCB,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(rcv),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetAttributeByName(he HELEMENT, name LPCSTR, value LPCWSTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetAttributeByName,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(value)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterClearAttributes(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterClearAttributes,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementIndex(he HELEMENT, p_index LPUINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementIndex,
		uintptr(he),
		uintptr(unsafe.Pointer(p_index)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementType(he HELEMENT, p_type *LPCSTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementType,
		uintptr(he),
		uintptr(unsafe.Pointer(p_type)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementTypeCB(he HELEMENT, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementTypeCB,
		uintptr(he),
		uintptr(unsafe.Pointer(rcv)),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetStyleAttributeCB(he HELEMENT, name LPCSTR, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetStyleAttributeCB,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(rcv),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetStyleAttribute(he HELEMENT, name LPCSTR, value LPCWSTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetStyleAttribute,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(value)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementLocation(he HELEMENT, p_location *Rect) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementLocation,
		uintptr(he),
		uintptr(unsafe.Pointer(p_location)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterScrollToView(he HELEMENT, SciterScrollFlags UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterScrollToView,
		uintptr(he),
		uintptr(SciterScrollFlags),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterUpdateElement(he HELEMENT, andForceRender SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterUpdateElement,
		uintptr(he),
		uintptr(andForceRender),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetCapture(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetCapture,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterReleaseCapture(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterReleaseCapture,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementHwnd(he HELEMENT, p_hwnd *HWINDOW, rootWindow SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementHwnd,
		uintptr(he),
		uintptr(unsafe.Pointer(p_hwnd)),
		uintptr(rootWindow),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCombineURL(he HELEMENT, szUrlBuffer LPWSTR, UrlBufferSize UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCombineURL,
		uintptr(he),
		uintptr(unsafe.Pointer(szUrlBuffer)),
		uintptr(UrlBufferSize),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSelectElements(he HELEMENT, CSS_selectors LPCSTR, callback uintptr, param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSelectElements,
		uintptr(he),
		uintptr(unsafe.Pointer(CSS_selectors)),
		uintptr(callback),
		uintptr(unsafe.Pointer(param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSelectElementsW(he HELEMENT, CSS_selectors LPCWSTR, callback uintptr, param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSelectElementsW,
		uintptr(he),
		uintptr(unsafe.Pointer(CSS_selectors)),
		uintptr(callback),
		uintptr(unsafe.Pointer(param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSelectParent(he HELEMENT, selector LPCSTR, depth UINT, heFound *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSelectParent,
		uintptr(he),
		uintptr(unsafe.Pointer(selector)),
		uintptr(depth),
		uintptr(unsafe.Pointer(heFound)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSelectParentW(he HELEMENT, selector LPCWSTR, depth UINT, heFound *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSelectParentW,
		uintptr(he),
		uintptr(unsafe.Pointer(selector)),
		uintptr(depth),
		uintptr(unsafe.Pointer(heFound)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetElementHtml(he HELEMENT /* const */, html *BYTE, htmlLength UINT, where UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetElementHtml,
		uintptr(he),
		uintptr(unsafe.Pointer(html)),
		uintptr(htmlLength),
		uintptr(where),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementUID(he HELEMENT, puid *UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementUID,
		uintptr(he),
		uintptr(unsafe.Pointer(puid)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementByUID(hwnd HWINDOW, uid UINT, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementByUID,
		uintptr(hwnd),
		uintptr(uid),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterShowPopup(hePopup HELEMENT, heAnchor HELEMENT, placement UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterShowPopup,
		uintptr(hePopup),
		uintptr(heAnchor),
		uintptr(placement),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterHidePopup(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterHidePopup,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementState(he HELEMENT, pstateBits *UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementState,
		uintptr(he),
		uintptr(unsafe.Pointer(pstateBits)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetElementState(he HELEMENT, stateBitsToSet UINT, stateBitsToClear UINT, updateView SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetElementState,
		uintptr(he),
		uintptr(stateBitsToSet),
		uintptr(stateBitsToClear),
		uintptr(updateView),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCreateElement(tagname LPCSTR, textOrNull LPCWSTR, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateElement,
		uintptr(unsafe.Pointer(tagname)),
		uintptr(unsafe.Pointer(textOrNull)),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCloneElement(he HELEMENT, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCloneElement,
		uintptr(he),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterInsertElement(he HELEMENT, hparent HELEMENT, index UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterInsertElement,
		uintptr(he),
		uintptr(hparent),
		uintptr(index),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterDetachElement(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDetachElement,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterDeleteElement(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDeleteElement,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetTimer(he HELEMENT, milliseconds UINT, timer_id UINT_PTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetTimer,
		uintptr(he),
		uintptr(milliseconds),
		uintptr(timer_id),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterDetachEventHandler(he HELEMENT, pep uintptr, tag LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterDetachEventHandler,
		uintptr(he),
		uintptr(pep),
		uintptr(unsafe.Pointer(tag)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterAttachEventHandler(he HELEMENT, pep uintptr, tag LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterAttachEventHandler,
		uintptr(he),
		uintptr(pep),
		uintptr(unsafe.Pointer(tag)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterWindowAttachEventHandler(hwndLayout HWINDOW, pep uintptr, tag LPVOID, subscription UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterWindowAttachEventHandler,
		uintptr(hwndLayout),
		uintptr(pep),
		uintptr(unsafe.Pointer(tag)),
		uintptr(subscription),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterWindowDetachEventHandler(hwndLayout HWINDOW, pep uintptr, tag LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterWindowDetachEventHandler,
		uintptr(hwndLayout),
		uintptr(pep),
		uintptr(unsafe.Pointer(tag)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSendEvent(he HELEMENT, appEventCode UINT, heSource HELEMENT, reason UINT_PTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSendEvent,
		uintptr(he),
		uintptr(appEventCode),
		uintptr(heSource),
		uintptr(reason),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterPostEvent(he HELEMENT, appEventCode UINT, heSource HELEMENT, reason UINT_PTR) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterPostEvent,
		uintptr(he),
		uintptr(appEventCode),
		uintptr(heSource),
		uintptr(reason),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCallBehaviorMethod(he HELEMENT, params *METHOD_PARAMS) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCallBehaviorMethod,
		uintptr(he),
		uintptr(unsafe.Pointer(params)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterRequestElementData(he HELEMENT, url LPCWSTR, dataType UINT, initiator HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterRequestElementData,
		uintptr(he),
		uintptr(unsafe.Pointer(url)),
		uintptr(dataType),
		uintptr(initiator),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterHttpRequest(
	he HELEMENT, // element to deliver data
	url LPCWSTR, // url
	dataType UINT, // data type, see SciterResourceType.
	requestType UINT, // one of REQUEST_TYPE values
	requestParams *REQUEST_PARAM, // parameters
	nParams UINT, // number of paramters
) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterHttpRequest,
		uintptr(he),
		uintptr(requestType),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetScrollInfo(he HELEMENT, scrollPos LPPOINT, viewRect *Rect, contentSize LPSIZE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetScrollInfo,
		uintptr(he),
		uintptr(unsafe.Pointer(scrollPos)),
		uintptr(unsafe.Pointer(viewRect)),
		uintptr(unsafe.Pointer(contentSize)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementIntrinsicWidths(he HELEMENT, pMinWidth *INT, pMaxWidth *INT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementIntrinsicWidths,
		uintptr(he),
		uintptr(unsafe.Pointer(pMinWidth)),
		uintptr(unsafe.Pointer(pMaxWidth)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementIntrinsicHeight(he HELEMENT, forWidth INT, pHeight *INT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementIntrinsicHeight,
		uintptr(he),
		uintptr(forWidth),
		uintptr(unsafe.Pointer(pHeight)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterIsElementVisible(he HELEMENT, pVisible *SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterIsElementVisible,
		uintptr(he),
		uintptr(unsafe.Pointer(pVisible)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterIsElementEnabled(he HELEMENT, pEnabled *SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterIsElementEnabled,
		uintptr(he),
		uintptr(unsafe.Pointer(pEnabled)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSortElements(he HELEMENT, firstIndex UINT, lastIndex UINT, cmpFunc uintptr, cmpFuncParam LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSortElements,
		uintptr(he),
		uintptr(firstIndex),
		uintptr(lastIndex),
		uintptr(unsafe.Pointer(cmpFunc)),
		uintptr(unsafe.Pointer(cmpFuncParam)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSwapElements(he1 HELEMENT, he2 HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSwapElements,
		uintptr(he1),
		uintptr(he2),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterTraverseUIEvent(evt UINT, eventCtlStruct LPVOID, bOutProcessed *SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterTraverseUIEvent,
		uintptr(evt),
		uintptr(unsafe.Pointer(eventCtlStruct)),
		uintptr(unsafe.Pointer(bOutProcessed)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCallScriptingMethod(he HELEMENT, name LPCSTR /* const */, argv *Value, argc UINT, retval *Value) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCallScriptingMethod,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(argv)),
		uintptr(argc),
		uintptr(unsafe.Pointer(retval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCallScriptingFunction(he HELEMENT, name LPCSTR /* const */, argv *Value, argc UINT, retval *Value) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCallScriptingFunction,
		uintptr(he),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(argv)),
		uintptr(argc),
		uintptr(unsafe.Pointer(retval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterEvalElementScript(he HELEMENT, script LPCWSTR, scriptLength UINT, retval *Value) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterEvalElementScript,
		uintptr(he),
		uintptr(unsafe.Pointer(script)),
		uintptr(scriptLength),
		uintptr(unsafe.Pointer(retval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterAttachHwndToElement(he HELEMENT, hwnd HWINDOW) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterAttachHwndToElement,
		uintptr(he),
		uintptr(hwnd),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterControlGetType(he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterControlGetType,
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetValue(he HELEMENT, pval *Value) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetValue,
		uintptr(he),
		uintptr(unsafe.Pointer(pval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetValue(he HELEMENT /* const */, pval *Value) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetValue,
		uintptr(he),
		uintptr(unsafe.Pointer(pval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetExpando(he HELEMENT, pval *Value, forceCreation SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetExpando,
		uintptr(he),
		uintptr(unsafe.Pointer(pval)),
		uintptr(forceCreation),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetObject(he HELEMENT, pval *interface{}, forceCreation SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetObject,
		uintptr(he),
		uintptr(unsafe.Pointer(pval)),
		uintptr(forceCreation),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetElementNamespace(he HELEMENT, pval *interface{}) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementNamespace,
		uintptr(he),
		uintptr(unsafe.Pointer(pval)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetHighlightedElement(hwnd HWINDOW, phe *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetHighlightedElement,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(phe)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetHighlightedElement(hwnd HWINDOW, he HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetHighlightedElement,
		uintptr(hwnd),
		uintptr(he),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeAddRef(hn HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeAddRef,
		uintptr(hn),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeRelease(hn HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeRelease,
		uintptr(hn),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeCastFromElement(he HELEMENT, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeCastFromElement,
		uintptr(he),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeCastToElement(hn HNODE, he *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeCastToElement,
		uintptr(hn),
		uintptr(unsafe.Pointer(he)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeFirstChild(hn HNODE, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeFirstChild,
		uintptr(hn),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeLastChild(hn HNODE, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeLastChild,
		uintptr(hn),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeNextSibling(hn HNODE, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeNextSibling,
		uintptr(hn),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodePrevSibling(hn HNODE, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodePrevSibling,
		uintptr(hn),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeParent(hnode HNODE, pheParent *HELEMENT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeParent,
		uintptr(hnode),
		uintptr(unsafe.Pointer(pheParent)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeNthChild(hnode HNODE, n UINT, phn *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeNthChild,
		uintptr(hnode),
		uintptr(n),
		uintptr(unsafe.Pointer(phn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeChildrenCount(hnode HNODE, pn *UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeChildrenCount,
		uintptr(hnode),
		uintptr(unsafe.Pointer(pn)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeType(hnode HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeType,
		uintptr(hnode),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeGetText(hnode HNODE, rcv uintptr, rcv_param LPVOID) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeGetText,
		uintptr(hnode),
		uintptr(rcv),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeSetText(hnode HNODE, text LPCWSTR, textLength UINT) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeSetText,
		uintptr(hnode),
		uintptr(unsafe.Pointer(text)),
		uintptr(textLength),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeInsert(hnode HNODE, what HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeInsert,
		uintptr(hnode),
		uintptr(what),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterNodeRemove(hnode HNODE, finalize SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeRemove,
		uintptr(hnode),
		uintptr(finalize),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCreateTextNode(text LPCWSTR, textLength UINT, phnode *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateTextNode,
		uintptr(unsafe.Pointer(text)),
		uintptr(textLength),
		uintptr(unsafe.Pointer(phnode)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterCreateCommentNode(text LPCWSTR, textLength UINT, phnode *HNODE) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateCommentNode,
		uintptr(unsafe.Pointer(text)),
		uintptr(textLength),
		uintptr(unsafe.Pointer(phnode)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) ValueInit(pval *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueInit,
		uintptr(unsafe.Pointer(pval)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueClear(pval *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueClear,
		uintptr(unsafe.Pointer(pval)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueCompare( /* const */ pval1 *Value /* const */, pval2 *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueCompare,
		uintptr(unsafe.Pointer(pval1)),
		uintptr(unsafe.Pointer(pval2)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueCopy(pdst *Value /* const */, psrc *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueCopy,
		uintptr(unsafe.Pointer(pdst)),
		uintptr(unsafe.Pointer(psrc)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueIsolate(pdst *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueIsolate,
		uintptr(unsafe.Pointer(pdst)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueType( /* const */ pval *Value, pType *UINT, pUnits *UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueType,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pUnits)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueStringData( /* const */ pval *Value, pChars *LPCWSTR, pNumChars *UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueStringData,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pChars)),
		uintptr(unsafe.Pointer(pNumChars)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueStringDataSet(pval *Value, chars LPCWSTR, numChars UINT, units UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueStringDataSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(chars)),
		uintptr(numChars),
		uintptr(units),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueIntData( /* const */ pval *Value, pData *INT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueIntData,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pData)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueIntDataSet(pval *Value, data INT, _type UINT, units UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueIntDataSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(data),
		uintptr(_type),
		uintptr(units),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueInt64Data( /* const */ pval *Value, pData *INT64) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueInt64Data,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pData)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueInt64DataSet(pval *Value, data INT64, _type UINT, units UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueInt64DataSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(data),
		uintptr(_type),
		uintptr(units),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueFloatData( /* const */ pval *Value, pData *FLOAT_VALUE) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueFloatData,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pData)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueFloatDataSet(pval *Value, data FLOAT_VALUE, _type UINT, units UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueFloatDataSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(data),
		uintptr(_type),
		uintptr(units),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueBinaryData( /* const */ pval *Value, pBytes *LPCBYTE, pnBytes *UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueBinaryData,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pBytes)),
		uintptr(unsafe.Pointer(pnBytes)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueBinaryDataSet(pval *Value, pBytes LPCBYTE, nBytes UINT, _type UINT, units UINT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueBinaryDataSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pBytes)),
		uintptr(nBytes),
		uintptr(_type),
		uintptr(units),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueElementsCount( /* const */ pval *Value, pn *INT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueElementsCount,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pn)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueNthElementValue( /* const */ pval *Value, n INT, pretval *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueNthElementValue,
		uintptr(unsafe.Pointer(pval)),
		uintptr(n),
		uintptr(unsafe.Pointer(pretval)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueNthElementValueSet(pval *Value, n INT /* const */, pval_to_set *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueNthElementValueSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(n),
		uintptr(unsafe.Pointer(pval_to_set)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueNthElementKey( /* const */ pval *Value, n INT, pretval *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueNthElementKey,
		uintptr(unsafe.Pointer(pval)),
		uintptr(n),
		uintptr(unsafe.Pointer(pretval)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueEnumElements( /* const */ pval *Value, penum uintptr, param LPVOID) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueEnumElements,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(penum)),
		uintptr(unsafe.Pointer(param)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueSetValueToKey(pval *Value /* const */, pkey *Value /* const */, pval_to_set *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueSetValueToKey,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pkey)),
		uintptr(unsafe.Pointer(pval_to_set)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueGetValueOfKey( /* const */ pval *Value /* const */, pkey *Value, pretval *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueGetValueOfKey,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pkey)),
		uintptr(unsafe.Pointer(pretval)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueToString(pval *Value, how VALUE_STRING_CVT_TYPE) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueToString,
		uintptr(unsafe.Pointer(pval)),
		uintptr(how),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueFromString(pval *Value, str LPCWSTR, strLength UINT, how VALUE_STRING_CVT_TYPE) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueFromString,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(str)),
		uintptr(strLength),
		uintptr(how),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueInvoke( /* const */ pval *Value, pthis *Value, argc UINT /* const */, argv *Value, pretval *Value, url LPCWSTR) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueInvoke,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pthis)),
		uintptr(argc),
		uintptr(unsafe.Pointer(argv)),
		uintptr(unsafe.Pointer(pretval)),
		uintptr(unsafe.Pointer(url)),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueNativeFunctorSet(pval *Value, pinvoke uintptr, prelease uintptr, tag LPVOID) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnValueNativeFunctorSet,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(pinvoke)),
		uintptr(unsafe.Pointer(prelease)),
		uintptr(tag),
	)
	return UINT(ret)
}

func (a *SciterAPI) ValueIsNativeFunctor( /* const */ pval *Value) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnValueIsNativeFunctor,
		uintptr(unsafe.Pointer(pval)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterOpenArchive(archiveData LPCBYTE, archiveDataLength UINT) HSARCHIVE {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterOpenArchive,
		uintptr(unsafe.Pointer(archiveData)),
		uintptr(archiveDataLength),
	)
	return HSARCHIVE(ret)
}

func (a *SciterAPI) SciterGetArchiveItem(harc HSARCHIVE, path LPCWSTR, pdata *LPCBYTE, pdataLength *UINT) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetArchiveItem,
		uintptr(harc),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(pdata)),
		uintptr(unsafe.Pointer(pdataLength)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterCloseArchive(harc HSARCHIVE) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCloseArchive,
		uintptr(harc),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterFireEvent( /* const */ evt *BEHAVIOR_EVENT_PARAMS, post SBOOL) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterFireEvent,
		uintptr(unsafe.Pointer(evt)),
		uintptr(post),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterGetCallbackParam(hwnd HWINDOW) LPVOID {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetCallbackParam,
		uintptr(hwnd),
	)
	return LPVOID(unsafe.Pointer(ret))
}

func (a *SciterAPI) SciterPostCallback(hwnd HWINDOW, wparam UINT_PTR, lparam UINT_PTR, timeoutms UINT) UINT_PTR {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterPostCallback,
		uintptr(hwnd),
		uintptr(wparam),
		uintptr(lparam),
		uintptr(timeoutms),
	)
	return UINT_PTR(ret)
}

func (a *SciterAPI) SciterCreateOnDirectXWindow(hwnd HWINDOW, pSwapChain *IUnknown) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterCreateOnDirectXWindow,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pSwapChain)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterRenderOnDirectXWindow(hwnd HWINDOW, elementToRenderOrNull HELEMENT, frontLayer SBOOL) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterRenderOnDirectXWindow,
		uintptr(hwnd),
		uintptr(elementToRenderOrNull),
		uintptr(frontLayer),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterRenderOnDirectXTexture(hwnd HWINDOW, elementToRenderOrNull HELEMENT, surface *IUnknown) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterRenderOnDirectXTexture,
		uintptr(hwnd),
		uintptr(elementToRenderOrNull),
		uintptr(unsafe.Pointer(surface)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterProcX(hwnd HWINDOW, pMsg *SCITER_X_MSG) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterProcX,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pMsg)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterAtomValue( /* const */ name *byte) UINT64 {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterAtomValue,
		uintptr(unsafe.Pointer(name)),
	)
	return UINT64(ret)
}

func (a *SciterAPI) SciterAtomNameCB(atomv UINT64, rcv *LPCSTR_RECEIVER, rcv_param LPVOID) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterAtomNameCB,
		uintptr(atomv),
		uintptr(unsafe.Pointer(rcv)),
		uintptr(unsafe.Pointer(rcv_param)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterSetGlobalAsset(pass *som_asset_t) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetGlobalAsset,
		uintptr(unsafe.Pointer(pass)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterGetElementAsset(el HELEMENT, nameAtom UINT64, ppass **som_asset_t) SCDOM_RESULT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetElementAsset,
		uintptr(el),
		uintptr(nameAtom),
		uintptr(unsafe.Pointer(ppass)),
	)
	return SCDOM_RESULT(ret)
}

func (a *SciterAPI) SciterSetVariable(hwndOrNull HWINDOW, name LPCSTR /* const */, pvalToSet *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterSetVariable,
		uintptr(hwndOrNull),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(pvalToSet)),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterGetVariable(hwndOrNull HWINDOW, name LPCSTR, pvalToGet *Value) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterGetVariable,
		uintptr(hwndOrNull),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(pvalToGet)),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterElementUnwrap( /* const */ pval *Value, ppElement *HELEMENT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterElementUnwrap,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(ppElement)),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterElementWrap(pval *Value, pElement HELEMENT) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterElementWrap,
		uintptr(unsafe.Pointer(pval)),
		uintptr(pElement),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterNodeUnwrap( /* const */ pval *Value, ppNode *HNODE) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeUnwrap,
		uintptr(unsafe.Pointer(pval)),
		uintptr(unsafe.Pointer(ppNode)),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterNodeWrap(pval *Value, pNode HNODE) UINT {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterNodeWrap,
		uintptr(unsafe.Pointer(pval)),
		uintptr(pNode),
	)
	return UINT(ret)
}

func (a *SciterAPI) SciterReleaseGlobalAsset(pass *som_asset_t) SBOOL {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterReleaseGlobalAsset,
		uintptr(unsafe.Pointer(pass)),
	)
	return SBOOL(ret)
}

func (a *SciterAPI) SciterExec(appCmd SCITER_APP_CMD, p1 UINT_PTR, p2 UINT_PTR) INT_PTR {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterExec,
		uintptr(appCmd),
		uintptr(p1),
		uintptr(p2),
	)
	return INT_PTR(ret)
}

func (a *SciterAPI) SciterWindowExec(hwnd HWINDOW, windowCmd UINT, p1 UINT_PTR, p2 UINT_PTR) INT_PTR {
	ret, _, _ := syscall.SyscallN(
		a.fnSciterWindowExec,
		uintptr(hwnd),
		uintptr(windowCmd),
		uintptr(p1),
		uintptr(p2),
	)
	return INT_PTR(ret)
}
