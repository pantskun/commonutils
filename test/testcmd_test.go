package main

import (
	"testing"
)

// func TestMain(t *testing.T) {
// 	cmd := exec.Command("testcmd.exe")

// 	var (
// 		stdout bytes.Buffer
// 		stdin  bytes.Buffer
// 	)

// 	cmd.Stdout = &stdout
// 	cmd.Stdin = &stdin

// 	stdin.WriteString("1\n")
// 	stdin.WriteString("12.66\n")

// 	// <-ch
// 	if err := cmd.Run(); err != nil {
// 		t.Fatal(err)
// 	}

// 	log.Println(stdout.String())
// }

type AnyType interface{}

type S struct {
	data string
}

func (s S) Val() string {
	return s.data
}

func (s *S) ChangeP(data string) {
	s.data = data
}

func (s S) ChangeV(data string) {
	s.data = data
}

type testStruct struct {
	handler func()
}

func addSlice(s interface{}, v interface{}, t *testing.T) {
	_, ok := (s).(*[]int)
	if !ok {
		t.Log("not *[]interface{}")
		return
	}

}

type InterfaceTest interface {
	Equal() bool
}

func TestMain(t *testing.T) {
	t0 := testStruct{}
	t1 := &t0
	t2 := &t0
	t.Log(t1 == t2)
}

func InterfaceFunc(i interface{}) {
	println(i)
}

func IsEqual(a, b interface{}) bool {
	return a == b
}
