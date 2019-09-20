// Code generated https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package multipletypesets

import "log"

type IntStringMap map[int]string

func (m IntStringMap) Has(key int) bool {
	_, ok := m[key]
	return ok
}

func (m IntStringMap) Get(key int) string {
	return m[key]
}

func (m IntStringMap) Set(key int, value string) IntStringMap {
	log.Println(value)
	m[key] = value
	return m
}

type Float64BoolMap map[float64]bool

func (m Float64BoolMap) Has(key float64) bool {
	_, ok := m[key]
	return ok
}

func (m Float64BoolMap) Get(key float64) bool {
	return m[key]
}

func (m Float64BoolMap) Set(key float64, value bool) Float64BoolMap {
	log.Println(value)
	m[key] = value
	return m
}
