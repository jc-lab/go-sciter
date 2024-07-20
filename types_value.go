package sciter

type SCITER_VALUE = Value

type NATIVE_FUNCTOR_INVOKE func(tag uintptr /* *void */, argc UINT /* const */, argv *Value, retval *Value) // retval may contain error definition
type NATIVE_FUNCTOR_RELEASE func(tag uintptr /* *void */)
