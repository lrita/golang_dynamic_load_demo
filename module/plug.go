package module

//#cgo LDFLAGS: -ldl -fPIC
//#include <dlfcn.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Module interface {
	Start() error
	Stop() error
}

func Plug(name, path string) error {

	cstring := C.CString(path)
	defer C.free(unsafe.Pointer(cstring))

	handler := C.dlopen(cstring, C.RTLD_LAZY|C.RTLD_LOCAL)
	if handler == nil {
		return fmt.Errorf("load %s ocurr %s", path, C.GoString(C.dlerror()))
	}

	return nil
}

func UnPlug(name string) error {
	// TODO close
	return nil
}
