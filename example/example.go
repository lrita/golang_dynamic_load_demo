package main

import "C"
import (
	"fmt"
	"math/rand"
	"time"
)

// "package main" and "func main() {}" must be defined in per module package
func main() {}

//export NewModule
func NewModule() interface{} {
	//This function must be defined in your module's package
	//It return your module which will be loads by golang_dynamic_load_demo
	return &exampleModule{
		id: rand.Int31(),
	}
}

func init() {
	fmt.Printf("run init() function...\n")
	rand.Seed(time.Now().Unix())
}

type exampleModule struct {
	id int32
}

func (m *exampleModule) Start() error {
	fmt.Printf("exampleModule Start...@ %d\n", m.id)
	return nil
}

func (m *exampleModule) Stop() error {
	fmt.Printf("exampleModule Stop...@ %d\n", m.id)
	return nil
}
