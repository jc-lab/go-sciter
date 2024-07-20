package sciter

type METHOD_PARAMS struct {
	MethodID UINT
}

type REQUEST_PARAM struct {
	Name  LPCWSTR
	Value LPCWSTR
}

// windows
type SciterWindowDelegate func(hwnd HWINDOW, msg UINT, wParam WPARAM, lParam LPARAM, pParam LPVOID, handled *SBOOL) LRESULT

type LPELEMENT_EVENT_PROC func(tag LPVOID, he HELEMENT, evtg UINT, prms LPVOID) SBOOL
type SciterElementCallback func(he HELEMENT, param LPVOID) SBOOL
type ELEMENT_COMPARATOR func(he1 HELEMENT, he2 HELEMENT, param LPVOID) INT

//typedef UINT SC_CALLBACK SciterHostCallback( LPSCITER_CALLBACK_NOTIFICATION pns, LPVOID callbackParam );

type LPCWSTR_RECEIVER func(str LPCWSTR, str_length UINT, param LPVOID)
type LPCSTR_RECEIVER func(str LPCSTR, str_length UINT, param LPVOID)
type LPCBYTE_RECEIVER func(str LPCBYTE, num_bytes UINT, param LPVOID)

type BEHAVIOR_EVENT_PARAMS struct {
	cmd      UINT     // BEHAVIOR_EVENTS
	heTarget HELEMENT // target element handler, in MENU_ITEM_CLICK this is owner element that caused this menu - e.g. context menu owner
	// In scripting this field named as Event.owner
	he     HELEMENT // source element e.g. in SELECTION_CHANGED it is new selected <option>, in MENU_ITEM_CLICK it is menu item (LI) element
	reason UINT_PTR // CLICK_REASON or EDIT_CHANGED_REASON - UI action causing change.
	// In case of custom event notifications this may be any
	// application specific value.
	data SCITER_VALUE // auxiliary data accompanied with the event. E.g. FORM_SUBMIT event is using this field to pass collection of values.

	name LPCWSTR // name of custom event (when cmd == CUSTOM)
}

type DATA_ARRIVED_PARAMS struct {
	initiator HELEMENT // element intiator of SciterRequestElementData request,
	data      LPCBYTE  // data buffer
	dataSize  UINT     // size of data
	dataType  UINT     // data type passed "as is" from SciterRequestElementData
	status    UINT     // status = 0 (dataSize == 0) - unknown error.
	// status = 100..505 - http response status, Note: 200 - OK!
	// status > 12000 - wininet error code, see ERROR_INTERNET_*** in wininet.h
	uri LPCWSTR // requested url
}
