package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = "implementation of mydocker"

	app.Commands = []cli.Command{
		RunCommand,
		InitCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

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
	//cmd := exec.Command(command)
	cmd := exec.Command("/proc/self/exe", "init", command)
	// 给self/exe进程增加NS隔离
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}

// 在self/exe进程中mount proc，并且启动用户进程并替换掉self/exe进程。使得用户进程获得self/exe的NS和PID等信息。
func Init(command string) {
	// TODO 改动(原理还不太懂): systemd 加入linux之后, mount namespace 就变成 shared by default, 所以你必须显示声明你要这个新的mount namespace独立。
	syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	cmd := exec.Command(command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Printf("Init Run() function err : %v\n", err)
		log.Fatal(err)
	}
}

var InitCommand = cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(c *cli.Context) error {
		command := c.Args().Get(0)
		Init(command)
		return nil
	},
}
