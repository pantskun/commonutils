package container

type Vector interface {
	Size() int
	Add(value Element)
	Insert(value Element, pos int)
	Find(value Element) int
	Remove(pos int)
}

type vector struct {
	values []Element
}

func NewVector() Vector {
	return &vector{}
}

func NewVectorWithSize(size int) Vector {
	return &vector{values: make([]Element, size)}
}

func (v *vector) Size() int {
	return len(v.values)
}

func (v *vector) Add(value Element) {
	v.values = append(v.values, value)
}

func (v *vector) Insert(value Element, pos int) {
	l := len(v.values)

	if pos >= l {
		v.values = append(v.values, value)
	} else {
		remain := v.values[pos+1:]
		v.values = append(v.values[:pos], value)
		v.values = append(v.values, remain...)
	}
}

func (v *vector) Remove(pos int) {
	l := len(v.values)

	if pos < l && pos >= 0 {
		v.values = append(v.values[:pos], (v.values[pos+1:])...)
	}
}

func (v *vector) Find(value Element) int {
	for i, v := range v.values {
		if Equal(v, value) {
			return i
		}
	}

	return -1
}
