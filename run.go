package main

import (
	"fmt"
	"go-docker/container"
)

func Run(cmd string, tty bool) error {
	parent := container.NewParentProcess(cmd, tty)
	if err := parent.Start(); err != nil {
		fmt.Println("Run:", err)
		return err
	}
	return parent.Wait()
}
