package module

import (
	"fmt"
	"sync"
)

var modules = make(map[string]Module)
var mu sync.RWMutex

func Register(name string, module Module) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := modules[name]; ok {
		return fmt.Errorf("duplicate module : %q", name)
	}
	modules[name] = module
	return nil
}

func UnRegister(name string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := modules[name]; !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	delete(modules, name)
	return nil
}

func Names() []string {
	names := make([]string, len(modules))
	mu.RLock()
	defer mu.RUnlock()
	for name := range modules {
		names = append(names, name)
	}
	return names
}

func Modules() []Module {
	ms := make([]Module, len(modules))
	mu.RLock()
	defer mu.RUnlock()
	for _, module := range modules {
		ms = append(ms, module)
	}
	return ms
}

func Start(name string) error {
	mu.RLock()
	defer mu.RUnlock()

	module, ok := modules[name]
	if !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	return module.Start()
}

func Stop(name string) error {
	mu.RLock()
	defer mu.RUnlock()

	module, ok := modules[name]
	if !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	return module.Stop()
}
