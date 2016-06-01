package main

import (
	"fmt"
	"time"

	"github.com/lrita/golang_dynamic_load_demo/module"
)

//This is a demo
func main() {

	err := module.Plug("example", "./lib/libexample.so")
	if err != nil {
		fmt.Printf("Load error: %v\n", err)
	}

	names := module.Names()
	fmt.Printf("Got %d modules.\n", len(names))

	for _, name := range names {
		if err := module.Start(name); err != nil {
			fmt.Printf("start %q module error : %v\n", name, err)
			continue
		}
		fmt.Printf("start %q module\n", name)
	}
	time.Sleep(time.Second) // mock do something
	for _, name := range names {
		if err := module.Stop(name); err != nil {
			fmt.Printf("stop %q module error : %v\n", name, err)
			continue
		}
		fmt.Printf("stop %q module\n", name)
	}
}
