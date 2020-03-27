package di

import (
	"reflect"
)

var simpleContainer = make(map[reflect.Type]interface{})

func SimpleSet(element interface{}) {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		t = v.Elem().Type()
	}

	simpleContainer[t] = element
}

func SimpleGet(element interface{}) interface{} {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		t = v.Elem().Type()
	}

	return simpleContainer[t]
}

func SimpleGetWithChecks(element interface{}) (interface{}, bool) {
	v := reflect.ValueOf(element)
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		t = v.Elem().Type()
	}

	e, ok := simpleContainer[t]

	return e, ok
}
