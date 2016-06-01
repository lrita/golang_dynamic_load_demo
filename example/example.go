package main

import "C"
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lrita/golang_dynamic_load_demo/module"
)

// "package main" and "func main() {}" must be defined in per module package
func main() {}

func init() {
	fmt.Printf("run init() function...\n")
	rand.Seed(time.Now().Unix())
	module.Register("example", &exampleModule{id: rand.Int31()})
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
