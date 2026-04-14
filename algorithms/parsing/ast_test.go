package parsing

import (
	"reflect"
	"testing"
)

func TestParseFunctionInfo(t *testing.T) {
	src := `package sample

import "fmt"

type Greeter struct{}

func Add(a, b int) int {
	return a + b
}

func helper() {}

func (g *Greeter) Hello(name string) {
	fmt.Println(name)
	helper()
}
`

	got, err := ParseFunctionInfo(src)
	if err != nil {
		t.Fatalf("ParseFunctionInfo returned error: %v", err)
	}

	want := []FunctionInfo{
		{
			Name:           "Add",
			Receiver:       "",
			ParameterCount: 2,
			ResultCount:    1,
			Calls:          nil,
		},
		{
			Name:           "helper",
			Receiver:       "",
			ParameterCount: 0,
			ResultCount:    0,
			Calls:          nil,
		},
		{
			Name:           "Hello",
			Receiver:       "*Greeter",
			ParameterCount: 1,
			ResultCount:    0,
			Calls:          []string{"fmt.Println", "helper"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected function info: got %#v want %#v", got, want)
	}
}
