package gobdd

import (
	"errors"
	"github.com/go-bdd/gobdd/context"
	"reflect"
)

func validateStepFunc(f interface{}) error {
	value := reflect.ValueOf(f)
	if value.Kind() != reflect.Func {
		return errors.New("the parameter should be a function")
	}

	if value.Type().NumOut() != 1 {
		return errors.New("the function should return only error")
	}
	val := value.Type().Out(0)
	errorInterface  := reflect.TypeOf((*error)(nil)).Elem()
	if !val.Implements(errorInterface) {
		return errors.New("the returned value should implement the Error interface")
	}

	if value.Type().NumIn() < 1 {
		return errors.New("the function should have Context as the first argument")
	}

	val = value.Type().In(0)
	n := val.ConvertibleTo(reflect.ValueOf(context.Context{}).Type())
	if !n {
		return errors.New("the returned value should implement the Error interface")
	}
	return nil
}
