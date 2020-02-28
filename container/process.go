package container

import (
	"os"
	"os/exec"
	"syscall"
)

// 创建一个会隔离namespace进程的Command
func NewParentProcess(command string, tty bool) *exec.Cmd {
	// ydd:copy当前go进程，返回一个命令，用于执行当前进程的init命令。
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...) // ydd: 启动另一个进程,另一个进程的可执行文件恰好是自己。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
