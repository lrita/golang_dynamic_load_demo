package module

//#cgo LDFLAGS: -ldl
//#include <dlfcn.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

//export Module
type Module interface {
	Start() error
	Stop() error
}

type NewModule func() interface{}

const defaultSymbol = "NewModule"

func Plug(name, path string) error {

	cstringPath := C.CString(path)
	defer C.free(unsafe.Pointer(cstringPath))
	cstringSymbol := C.CString(defaultSymbol)
	defer C.free(unsafe.Pointer(cstringSymbol))

	handler := C.dlopen(cstringPath, C.RTLD_LAZY|C.RTLD_LOCAL)
	if handler == nil {
		return fmt.Errorf("load %s ocurr %s", path, C.GoString(C.dlerror()))
	}

	symbol := C.dlsym(handler, cstringSymbol)
	if symbol == nil {
		C.dlclose(handler)
		return fmt.Errorf("load %s not found %", path, defaultSymbol)
	}

	interf := (*(*NewModule)(symbol))()
	if interf == nil {
		C.dlclose(handler)
		return fmt.Errorf("load %s get nil module", path)
	}

	module, ok := interf.(Module)
	if !ok {
		C.dlclose(handler)
		return fmt.Errorf("load %s get nil module", path)
	}

	if err := Register(name, module); err != nil {
		C.dlclose(handler)
		return fmt.Errorf("load %s register module : %v", path, err)
	}

	return nil
}

func UnPlug(name string) error {
	// TODO C.dlclose(handler)
	return nil
}
