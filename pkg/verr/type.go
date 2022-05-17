package verr

import "reflect"

type Type interface {
}

type Unknown struct {
	Type
}

func IsType(t, target Type) bool {
	return typesEqual(t, target)
}

func getTypeName(t Type) string {
	/*if t := reflect.TypeOf(t); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}*/
	return getElementType(t).Name()
}

func typesEqual(t1, t2 Type) bool {
	return getElementType(t1) == getElementType(t2)
}

func getElementType(t Type) reflect.Type {
	if t := reflect.TypeOf(t); t.Kind() == reflect.Ptr {
		return t.Elem()
	} else {
		return t
	}
}
