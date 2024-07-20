package sciter

type SCITER_CALLBACK_NOTIFICATION struct {
	Code UINT    /**< [in] one of the codes above.*/
	Hwnd HWINDOW /**< [in] HWINDOW of the window this callback was attached to.*/
}
type LPSCITER_CALLBACK_NOTIFICATION *SCITER_CALLBACK_NOTIFICATION

type SciterHostCallback func(pns LPSCITER_CALLBACK_NOTIFICATION, callbackParam LPVOID) UINT
type LPSciterHostCallback = SciterHostCallback

// SCITER_CREATE_WINDOW_FLAGS defines window creation flags
type SCITER_CREATE_WINDOW_FLAGS UINT

// SCITER_WINDOW_CMD defines window commands
type SCITER_WINDOW_CMD UINT

const (
	SCITER_WINDOW_SET_STATE              SCITER_WINDOW_CMD = 1  // p1 - SCITER_WINDOW_STATE, p2 - N/A
	SCITER_WINDOW_GET_STATE              SCITER_WINDOW_CMD = 2  // p1 - N/A , p2 - N/A, returns SCITER_WINDOW_STATE
	SCITER_WINDOW_ACTIVATE               SCITER_WINDOW_CMD = 3  // p1 - BOOL, true - bring_to_front , p2 - N/A
	SCITER_WINDOW_SET_PLACEMENT          SCITER_WINDOW_CMD = 4  // p1 - const POINT*, position, p2 const SIZE* - dimension, in ppx, either one can be null
	SCITER_WINDOW_GET_PLACEMENT          SCITER_WINDOW_CMD = 5  // p1 - POINT*, position, p2 SIZE* - dimension, in ppx, either one can be null
	SCITER_WINDOW_GET_VULKAN_ENVIRONMENT SCITER_WINDOW_CMD = 20 // p1 - &SciterVulkanEnvironment, p2 - sizeof(SciterVulkanEnvironment)
	SCITER_WINDOW_GET_VULKAN_CONTEXT     SCITER_WINDOW_CMD = 21 // p1 - &SciterVulkanContext, p2 - sizeof(SciterVulkanContext)
	SCITER_WINDOW_SET_VULKAN_BRIDGE      SCITER_WINDOW_CMD = 22 // p1 - SciterWindowVulkanBridge*, p2 - N/A
)

// SCITER_WINDOW_STATE defines window states
type SCITER_WINDOW_STATE UINT_PTR

const (
	SCITER_WINDOW_STATE_CLOSED      SCITER_WINDOW_STATE = 0 // close window
	SCITER_WINDOW_STATE_SHOWN       SCITER_WINDOW_STATE = 1
	SCITER_WINDOW_STATE_MINIMIZED   SCITER_WINDOW_STATE = 2
	SCITER_WINDOW_STATE_MAXIMIZED   SCITER_WINDOW_STATE = 3
	SCITER_WINDOW_STATE_HIDDEN      SCITER_WINDOW_STATE = 4
	SCITER_WINDOW_STATE_FULL_SCREEN SCITER_WINDOW_STATE = 5
)

type SCN_LOAD_DATA struct {
	code UINT    /**< [in] one of the codes above.*/
	hwnd HWINDOW /**< [in] HWINDOW of the window this callback was attached to.*/

	uri LPCWSTR /**< [in] Zero terminated string, fully qualified uri, for example "http://server/folder/file.ext".*/

	outData     LPCBYTE /**< [in,out] pointer to loaded data to return. if data exists in the cache then this field contain pointer to it*/
	outDataSize UINT    /**< [in,out] loaded data size to return.*/
	dataType    UINT    /**< [in] SciterResourceType */

	requestId HREQUEST /**< [in] request handle that can be used with sciter-x-request API */

	principal HELEMENT
	initiator HELEMENT
}

type SCN_ATTACH_BEHAVIOR struct {
	code UINT    /**< [in] one of the codes above.*/
	hwnd HWINDOW /**< [in] HWINDOW of the window this callback was attached to.*/

	element      HELEMENT /**< [in] target DOM element handle*/
	behaviorName LPCSTR   /**< [in] zero terminated string, string appears as value of CSS behavior:"???" attribute.*/

	elementProc uintptr /**< [out] pointer to ElementEventProc function.*/
	elementTag  LPVOID  /**< [out] tag value, passed as is into pointer ElementEventProc function.*/

}
