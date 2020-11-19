package main

var instance Single

type Single interface {
	testFunc() int
}

type single struct {
	data int
}

func GetInstance() Single {
	if instance == nil {
		newInstance()
	}

	return instance
}

func newInstance() Single {
	return &single{data: 0}
}

func (s single) testFunc() int {
	return s.data
}
