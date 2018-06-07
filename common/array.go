package common

import (
	"errors"
	"fmt"
	"reflect"
)

type Array []interface{}

func NewArray(s interface{}) (*Array, error) {
	switch reflect.TypeOf(s).Kind() {
	case reflect.Slice, reflect.Array:
		v := reflect.ValueOf(s)
		ret := make(Array, v.Len())

		for i := 0; i < v.Len(); i++ {
			ret[i] = v.Index(i).Interface()
		}

		return &ret, nil
	default:
		return nil, errors.New("s is not array or slice")
	}
}

func (a *Array) miniArrayMap(f func(e interface{}) (interface{}, error)) (*Array, error) {
	ret := make(Array, len(*a))

	for i, e := range *a {
		var err error
		ret[i], err = f(e)
		if err != nil {
			return nil, err
		}
	}

	return &ret, nil
}

func (a *Array) Map(f func(e interface{}) (interface{}, error)) (*Array, error) {
	return a.miniArrayMap(f)
}

func (a *Array) GroupField(f string) (map[interface{}][]interface{}, error) {
	ret := map[interface{}][]interface{}{}

	for _, e := range *a {
		v := reflect.ValueOf(e)
		switch v.Kind() {
		case reflect.Struct:
			switch v.FieldByName(f).Kind() {
			case reflect.String,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64:
				ret[v.FieldByName(f).Interface()] = append(ret[v.FieldByName(f).Interface()], e)
			default:
				return nil, fmt.Errorf("unsupported type: %v", v.Kind())
			}
		default:
			return nil, fmt.Errorf("unsupported type: %v", v.Kind())
		}
	}

	return ret, nil
}
