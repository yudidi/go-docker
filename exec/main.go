package main

import (
	"log"
	"os"
	"syscall"
)

// syscall.Exec启动子进程会吃掉父进程，并占据pid等全部信息。(本质是替换进程镜像)
func main() {
	log.Printf("pid:%d\n", os.Getpid())
	command := "/bin/sh"
	if err := syscall.Exec(command, []string{}, os.Environ()); err != nil {
		log.Printf("syscall.Exec err: %v\n", err)
		log.Fatal(err)
	}
}
