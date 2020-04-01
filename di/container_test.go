package di

import (
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
	container = make(map[reflect.Type]reflect.Value)
}

// This function is performed AFTER all tests in this package
func tearDown() {
	// it's only for testing
	container = nil
}

func TestSet_Success(t *testing.T) {
	err := Set(&simpleType{})

	if err != nil {
		t.Error("unable to set pointer to struct")
	}
}

func TestSet_ArgumentMustBePointer(t *testing.T) {
	err := Set(simpleType{})

	if err != ErrMustBePointer {
		t.Error("set returns wrong error message for value by reference")
	}
}

func TestSet_ArgumentMustBeStruct(t *testing.T) {
	value := "custom message"
	err := Set(&value)

	if err != ErrMustBeStructPointer {
		t.Error("set returns wrong error message for value by pointer")
	}
}
