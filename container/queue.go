package container

type Queue struct {
	values []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewQueueWithSize(size int) *Queue {
	return &Queue{values: make([]interface{}, size)}
}

func (q *Queue) Push(value interface{}) {
	q.values = append(q.values, value)
}

func (q *Queue) Pop() interface{} {
	l := len(q.values)
	if l == 0 {
		return nil
	}

	t := q.values[0]
	q.values = q.values[1:]

	return t
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}
