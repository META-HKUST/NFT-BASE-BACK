package utils

import (
	"reflect"
	"unsafe"
)

func CheckEmpty(s string) bool {
	if s == "" {
		return false
	} else {
		return true
	}
}

func CheckIntEmpty(s int) bool {
	if s == 0 {
		return false
	} else {
		return true
	}
}

func CheckAnyEmpty(ss []string) bool {
	for _, i := range ss {
		if i == "" {
			return false
		}
	}
	return true
}

func CheckEmptyStruct(st interface{}) bool {

	val := reflect.ValueOf(st).Elem()
	typ := reflect.TypeOf(st).Elem()

	for i := 0; i < typ.NumField(); i++ {

		if typ.Field(i).Tag.Get("must") == "true" && val.Field(i).IsZero() {
			return false
		}
	}
	return true
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
