// Code generated with https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package math

func Float32Max(fn func(a, b float32) bool, a, b float32) float32 {
	if fn(a, b) {
		return a
	}
	return b
}

func Float64Max(fn func(a, b float64) bool, a, b float64) float64 {
	if fn(a, b) {
		return a
	}
	return b
}

func IntMax(fn func(a, b int) bool, a, b int) int {
	if fn(a, b) {
		return a
	}
	return b
}

func Int16Max(fn func(a, b int16) bool, a, b int16) int16 {
	if fn(a, b) {
		return a
	}
	return b
}

func Int32Max(fn func(a, b int32) bool, a, b int32) int32 {
	if fn(a, b) {
		return a
	}
	return b
}

func Int64Max(fn func(a, b int64) bool, a, b int64) int64 {
	if fn(a, b) {
		return a
	}
	return b
}

func Int8Max(fn func(a, b int8) bool, a, b int8) int8 {
	if fn(a, b) {
		return a
	}
	return b
}

func UintMax(fn func(a, b uint) bool, a, b uint) uint {
	if fn(a, b) {
		return a
	}
	return b
}

func Uint16Max(fn func(a, b uint16) bool, a, b uint16) uint16 {
	if fn(a, b) {
		return a
	}
	return b
}

func Uint32Max(fn func(a, b uint32) bool, a, b uint32) uint32 {
	if fn(a, b) {
		return a
	}
	return b
}

func Uint64Max(fn func(a, b uint64) bool, a, b uint64) uint64 {
	if fn(a, b) {
		return a
	}
	return b
}

func Uint8Max(fn func(a, b uint8) bool, a, b uint8) uint8 {
	if fn(a, b) {
		return a
	}
	return b
}
