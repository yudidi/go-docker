package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go-docker/container"
	"time"
)

// 创建namespace隔离的容器进程
// 启动容器
var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroups limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container args")
		}
		// cmd 为容器启动后运行的第一个命令程序
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		return Run(cmd, tty)
	},
}

// 初始化容器内容,挂载proc文件系统，运行用户执行程序
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		cmd := context.Args().Get(0)
		fmt.Println(cmd)
		for {
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
		return nil
		//return container.RunContainerInitProcess(cmd, nil)
	},
}
