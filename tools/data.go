package tools

//Data is an interface for test sort.
type Data interface {
	Add(i interface{})
	Get() []interface{}
}

//In for test int sort
type In struct {
	Elem []interface{}
}

//Add for append test data
func (i *In) Add(n interface{}) {
	i.Elem = append(i.Elem, n)
}

//Get for get test data
func (i *In) Get() []interface{} {
	return i.Elem
}
