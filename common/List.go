package common

type Any = interface{}

type List interface {
	Get(index int) Any
	Set(index int, elem Any)
	Slice(begin, end int) List
	Length() int
	Append(elem Any) List
	ToSlice() []Any
}

type SimpleList struct {
	list []Any
}

func FromSlice(slc []Any) *SimpleList {
	return &SimpleList{list: slc}
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

func (l *SimpleList) Append(elem Any) List {
	l.list = append(l.list, elem)
	return l
}

func (l *SimpleList) ToSlice() []Any {
	return l.list
}
