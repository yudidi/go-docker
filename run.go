package main

import (
	"fmt"
	"go-docker/container"
)

func Run(cmd string, tty bool) error {
	// ydd:返回一个命令，该命令可 使用与当前go进程相同的可执行文件，初始化进程本身。 TODO 那这2个进程的pid一样吗？
	parent := container.NewParentProcess(cmd, tty) // ydd: parent是cmd(/bin/sh，容器内第一个运行的进程)的父进程
	if err := parent.Start(); err != nil {         // ydd: 执行容器进程的init方法
		fmt.Println("Run:", err)
		return err
	}
	return parent.Wait()
}
