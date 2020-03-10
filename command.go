package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name: "run",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(c *cli.Context) error {
		tty := c.Bool("it")
		command := c.Args().Get(0)
		return Run(command, tty)
	},
}

func Run(command string, tty bool) error {
	cmd := exec.Command(command)
	// for kinds of namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}
	cmd.Run()
}
