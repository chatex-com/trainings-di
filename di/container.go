package di

import (
	"errors"
	"reflect"
)

var (
	container = make(map[reflect.Type]reflect.Value)

	ErrMustBePointer                  = errors.New("element must be a pointer")
	ErrMustBeStructPointer            = errors.New("element must be a pointer to struct")
	ErrMustBePointerToPointer         = errors.New("element must be a pointer to pointer")
	ErrMustBePointerToPointerToStruct = errors.New("element must be a pointer to pointer to struct (**element)")
	ErrElementNotFound                = errors.New("unable to find element")
)

func Set(element interface{}) error {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() != reflect.Ptr {
		return ErrMustBePointer
	}

	t = v.Elem().Type()

	if t.Kind() != reflect.Struct {
		return ErrMustBeStructPointer
	}

	container[t] = v

	return nil
}

func LoadPtr(element interface{}) error {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() != reflect.Ptr {
		return ErrMustBePointer
	}

	v = v.Elem()
	t = v.Type()
	if t.Kind() != reflect.Struct {
		return ErrMustBeStructPointer
	}

	item, ok := container[t]
	if !ok {
		return ErrElementNotFound
	}

	v.Set(item.Elem())

	return nil
}

func LoadPtrToPtr(element interface{}) error {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() != reflect.Ptr {
		return ErrMustBePointer
	}

	v = v.Elem()
	t = v.Type()

	if t.Kind() != reflect.Ptr {
		return ErrMustBePointerToPointer
	}

	t = v.Elem().Type()
	if t.Kind() != reflect.Struct {
		return ErrMustBePointerToPointerToStruct
	}

	item, ok := container[t]
	if !ok {
		return ErrElementNotFound
	}

	v.Set(item.Elem().Addr())

	return nil
}
