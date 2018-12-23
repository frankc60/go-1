package main

import (
		"fmt"
		"os/exec"
		"log"
)

func main() {
	fmt.Println("hello world, dir listing")

	cmd := exec.Command("sh", "-c", "ls -a; echo 1>&2")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
	
}

//execute: go run main.go

//to compile: go build main.go       , this creates "main"
//then ./main    (executable), 1 binary