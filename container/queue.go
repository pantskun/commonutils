package container

type Queue interface {
	Size() int
	IsEmpty() bool
	Push(value Element)
	Pop() Element
}

type queue struct {
	values []Element
}

func NewQueue() Queue {
	return &queue{}
}

func NewQueueWithSize(size int) Queue {
	return &queue{values: make([]Element, size)}
}

func (q *queue) Size() int {
	return len(q.values)
}

func (q *queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *queue) Push(value Element) {
	q.values = append(q.values, value)
}

func (q *queue) Pop() Element {
	l := len(q.values)
	if l == 0 {
		return nil
	}

	t := q.values[0]
	q.values = q.values[1:]

	return t
}
