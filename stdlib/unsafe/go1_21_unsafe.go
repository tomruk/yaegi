// Code generated by 'yaegi extract unsafe'. DO NOT EDIT.

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package unsafe

import (
	"reflect"
	"unsafe"
)

func init() {
	Symbols["unsafe/unsafe"] = map[string]reflect.Value{
		// type definitions
		"Pointer": reflect.ValueOf((*unsafe.Pointer)(nil)),
	}
}
