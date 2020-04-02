package di

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type simpleType struct{}

// Sometimes necessary for a test program to do extra setup or teardown before or after testing
// For our case we should recreate container for idempotent tests
func TestMain(m *testing.M) {
	setUp()
	exitCode := m.Run()
	tearDown()
	os.Exit(exitCode)
}

// This function is performed BEFORE all tests in this package
func setUp() {
	fmt.Println("it's just a message BEFORE all tests")
}

// This function is performed AFTER all tests in this package
func tearDown() {
	fmt.Println("it's just a message AFTER all tests")
}

func refreshInternalContainer() {
	container = make(map[reflect.Type]reflect.Value)
}

func TestSet_Success(t *testing.T) {
	refreshInternalContainer()

	err := Set(&simpleType{})

	if err != nil {
		t.Error("unable to set pointer to struct")
	}
}

func TestSet_ArgumentMustBePointer(t *testing.T) {
	refreshInternalContainer()

	err := Set(simpleType{})

	if err != ErrMustBePointer {
		t.Error("set returns wrong error message for value by reference")
	}
}

func TestSet_ArgumentMustBeStruct(t *testing.T) {
	refreshInternalContainer()

	value := "custom message"
	err := Set(&value)

	if err != ErrMustBeStructPointer {
		t.Error("set returns wrong error message for value by pointer")
	}
}

func TestLoadPtr_ElementNotFound(t *testing.T) {
	refreshInternalContainer()

	err := LoadPtr(&simpleType{})

	if err != ErrElementNotFound {
		t.Error("load returns wrong error message for element that does not exist")
	}
}

func TestLoadPtr_ArgumentShouldBePointer(t *testing.T) {
	refreshInternalContainer()

	err := LoadPtr(simpleType{})

	if err != ErrMustBePointer {
		t.Error("load returns wrong error message for element by value")
	}
}

func TestLoadPtr_ArgumentShouldBeStruct(t *testing.T) {
	refreshInternalContainer()

	element := "just a string"
	err := LoadPtr(&element)

	if err != ErrMustBeStructPointer {
		t.Error("load returns wrong error message for non struct element")
	}
}

func TestLoadPtr_Success(t *testing.T) {
	refreshInternalContainer()

	original := &simpleType{}
	_ = Set(original)

	target := simpleType{}
	err := LoadPtr(&target)
	if err != nil {
		t.Error("element should be loaded")
	}
}
