package main

import (
	"bytes"
	"log"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("testcmd.exe")

	var (
		stdout bytes.Buffer
		stdin  bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stdin = &stdin

	// ch := make(chan int)

	// go func() {
	// 	log.Println("start cmd")

	// 	if err := cmd.Run(); err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	log.Println("end cmd")

	// 	time.Sleep(5 * time.Second)

	// 	ch <- 0
	// }()

	stdin.WriteString("1\n")
	stdin.WriteString("12.66\n")

	// <-ch
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	log.Println(stdout.String())
}
