// Code generated by 'yaegi extract image/color/palette'. DO NOT EDIT.

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"image/color/palette"
	"reflect"
)

func init() {
	Symbols["image/color/palette/palette"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Plan9":   reflect.ValueOf(&palette.Plan9).Elem(),
		"WebSafe": reflect.ValueOf(&palette.WebSafe).Elem(),
	}
}
