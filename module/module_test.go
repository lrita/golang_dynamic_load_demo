package module

import (
	"strings"
	"testing"
)

type mockModule struct{}

func (m *mockModule) Start() error { return nil }
func (m *mockModule) Stop() error  { return nil }

func TestDuplicateRegister(t *testing.T) {
	module := &mockModule{}

	err := Register("testing", module)
	if err != nil {
		t.Errorf("error ocurr : %v", err)
	}
	err = Register("testing", module)
	if err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("no error ocurr, want a \"duplicate\" error.")
	}
}
