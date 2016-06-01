package module

import (
	"fmt"
	"sync"
)

var Modules = make(map[string]Module)
var mu sync.RWMutex

func Register(name string, module Module) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := Modules[name]; ok {
		return fmt.Errorf("duplicate module : %q", name)
	}
	Modules[name] = module
	fmt.Printf("Modules @ %p\n", Modules)
	return nil
}

func UnRegister(name string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := Modules[name]; !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	delete(Modules, name)
	return nil
}

func Names() []string {
	names := make([]string, len(Modules))
	mu.RLock()
	defer mu.RUnlock()
	for name := range Modules {
		names = append(names, name)
	}
	fmt.Printf("Modules @ %p\n", Modules)
	return names
}

func GetModules() []Module {
	ms := make([]Module, len(Modules))
	mu.RLock()
	defer mu.RUnlock()
	for _, module := range Modules {
		ms = append(ms, module)
	}
	return ms
}

func Start(name string) error {
	mu.RLock()
	defer mu.RUnlock()

	module, ok := Modules[name]
	if !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	return module.Start()
}

func Stop(name string) error {
	mu.RLock()
	defer mu.RUnlock()

	module, ok := Modules[name]
	if !ok {
		return fmt.Errorf("not found module : %q", name)
	}
	return module.Stop()
}
