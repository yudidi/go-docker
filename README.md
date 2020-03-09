# init命令
2个功能
1. 在容器中启动用户传入的进程，用户传入的进程会成为ns中第一个启动的进程。
2. 重新挂载proc,使得proc只显示容器的统计和运行信息？TODO



# syscall.Exec启动进程和os/exec.Command启动进程的区别

* os/exec.Command("/bin/sh")启动/bin/sh程序
可以看到父子进程的PID是不同的

```
func main()  {
    log.Printf("pid:%d\n", os.Getpid())
    cmd := exec.Command("/bin/sh")
    
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
    }
}
// 在shell中运行该程序会启动3个进程
[root@192 invoke_other]# go run cmd/main.go
2020/03/08 06:13:14 pid:18266
sh-4.2# echo $$
18269

// ps -ef //查看到启动了3个进程,依次是shell进程,go进程,sh进程
root      18244  18128  1 06:13 pts/0    00:00:00 go run cmd/main.go
root      18266  18244  0 06:13 pts/0    00:00:00 /tmp/go-build708334208/b001/exe/main
root      18269  18266  0 06:13 pts/0    00:00:00 sh
```

* syscall.Exec启动/bin/sh进程
可以看到父子进程的PID是相同的，`也就是说子进程已经吃掉了父进程`,pid18318已经是/bin/sh进程了，不再是go进程了。

```
func main()  {
    log.Printf("pid:%d\n", os.Getpid())
    command := "/bin/sh"
    if err := syscall.Exec(command, []string{command}, os.Environ()); err != nil {
    }   
}
// 在shell中运行该程序会启动2个进程
[root@192 invoke_other]# go run exec/main.go
2020/03/08 06:17:37 pid:18318
[root@192 invoke_other]# echo $$
18318
// 可以看到go进程已经没有了
root      18296  18128  1 06:17 pts/0    00:00:00 go run exec/main.go
root      18318  18296  0 06:17 pts/0    00:00:00 [sh]
root      18342  18197  0 06:17 pts/1    00:00:00 ps -ef
```

 
 # 代码改动
 ```
 // 启动init进程
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		fmt.Printf("init PID: %d \n", os.Getpid())
		cmd := context.Args().Get(0)
		return container.RunContainerInitProcess(cmd, nil)
	},
}

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("用户期望容器执行的第一个command %s", command)
	// 挂载
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		return err
	}

	err = syscall.Exec(command, []string{command}, os.Environ())
	if err != nil {
		fmt.Println("RunContainerInitProcess:", err)
		return err
	}
	return nil
}
 ```
 
 # 编译和执行结果如下
 ```
[root@192 go-docker]# go build .
[root@192 go-docker]# ./go-docker run --ti /bin/sh
run PID: 16952
init PID: 1 // 从这里可以看到/bin/sh进程id已经被欺骗为1了。
{"level":"info","msg":"用户期望容器执行的第一个command /bin/sh","time":"2020-03-08T03:32:32-04:00"}
sh-4.2#
 ```
