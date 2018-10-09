package common

import (
	"fmt"
	"reflect"
)

var (
	IsNotSliceError = fmt.Errorf("The type of value is not slice. ")
)

type Any = interface{}

type List interface {
	Get(index int) Any
	Set(index int, elem Any)
	Slice(begin, end int) List
	Length() int
	Append(elems ... Any) List
	ToSlice() []Any
}

type SimpleList struct {
	list []Any
}

func SliceToSimpleList(slc Any) *SimpleList {
	if reflect.TypeOf(slc).Kind() != reflect.Slice {
		panic(IsNotSliceError)
	}
	v := reflect.ValueOf(slc)
	list := make([]Any, v.Len())
	for i := 0; i < v.Len(); i++ {
		list[i] = v.Index(i).Interface()
	}
	return &SimpleList{list}
}

func (l *SimpleList) Get(index int) Any {
	return l.list[index]
}

func (l *SimpleList) Set(index int, elem Any) {
	l.list[index] = elem
}

func (l *SimpleList) Slice(begin, end int) List {
	return &SimpleList{
		list: l.list[begin:end],
	}
}

func (l *SimpleList) Length() int {
	return len(l.list)
}

func (l *SimpleList) Append(elems ... Any) List {
	l.list = append(l.list, elems...)
	return l
}

func (l *SimpleList) ToSlice() []Any {
	return l.list
}
