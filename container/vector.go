package container

type Vector struct {
	values []Element
}

func (v *Vector) InitWithSize(size int) {
	v.values = make([]Element, size)
}

func (v *Vector) Size() int {
	return len(v.values)
}

func (v *Vector) Add(value Element) {
	v.values = append(v.values, value)
}

func (v *Vector) Insert(value Element, pos int) {
	l := len(v.values)

	if pos >= l {
		v.values = append(v.values, value)
	} else {
		remain := v.values[pos+1:]
		v.values = append(v.values[:pos], value)
		v.values = append(v.values, remain...)
	}
}

func (v *Vector) Get(pos int) Element {
	return v.values[pos]
}

func (v *Vector) Remove(pos int) {
	l := len(v.values)

	if pos < l && pos >= 0 {
		v.values = append(v.values[:pos], (v.values[pos+1:])...)
	}
}

func (v *Vector) Find(value Element, equal func(Element, Element) bool) int {
	for i, v := range v.values {
		if equal(v, value) {
			return i
		}
	}

	return -1
}
