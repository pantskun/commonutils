package container

type Queue struct {
	values []Element
}

func (q *Queue) InitWithSize(size int) {
	q.values = make([]Element, size)
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) Push(value Element) {
	q.values = append(q.values, value)
}

func (q *Queue) Pop() Element {
	l := len(q.values)
	if l == 0 {
		return nil
	}

	t := q.values[0]
	q.values = q.values[1:]

	return t
}

func (q *Queue) Has(value Element, equal func(Element, Element) bool) bool {
	for _, v := range q.values {
		if equal(v, value) {
			return true
		}
	}

	return false
}
