package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Printf("pid:%d\n", os.Getpid())
	cmd := exec.Command("sh")

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("Init Run() function err : %v\n", err)
		log.Fatal(err)
	}

}
