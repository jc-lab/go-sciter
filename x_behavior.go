package sciter

type SCRIPTING_METHOD_PARAMS struct {
	name                          LPCSTR ///< method name
	argv/* const */ *SCITER_VALUE ///< vector of arguments
	argc                          UINT         ///< argument count
	result                        SCITER_VALUE ///< return value
}
