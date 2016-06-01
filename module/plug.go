package module

import (
	"fmt"

	"github.com/rainycape/dl"
)

type Module interface {
	Start() error
	Stop() error
}

type NewModule func() interface{}

const defaultSymbol = "NewModule"

func Plug(name, path string) error {

	dlHandler, err := dl.Open(path, dl.RTLD_LAZY|dl.RTLD_LOCAL)
	if err != nil {
		return err
	}

	var handlerFunc NewModule
	err = dlHandler.Sym(defaultSymbol, &handlerFunc)
	if err != nil {
		dlHandler.Close()
		return err
	}

	moduleInterface := handlerFunc()
	if moduleInterface == nil {
		dlHandler.Close()
		return fmt.Errorf("nil interface")
	}

	module, ok := moduleInterface.(Module)
	if !ok {
		dlHandler.Close()
		return fmt.Errorf("load %s get nil module", path)
	}

	if err := Register(name, module); err != nil {
		dlHandler.Close()
		return fmt.Errorf("load %s register module : %v", path, err)
	}

	return nil
}

func UnPlug(name string) error {
	// TODO close
	return nil
}
