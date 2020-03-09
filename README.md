# init命令
2个功能
1. 在容器中启动用户传入的进程，用户传入的进程会成为ns中第一个启动的进程。
2. 重新挂载proc,使得proc只显示容器的统计和运行信息？TODO



# syscall.Exec启动进程和os/exec.Command启动进程的区别

* os/exec.Command("/bin/sh")启动/bin/sh程序

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
// output // 可以看到PID是不同的
root@nicktming:~/go/src/github.com/nicktming/mydocker/test/syscall# go run TestExec.go 
2019/03/25 23:47:29 pid:20255
# echo $$
20258
```

* syscall.Exec启动/bin/sh进程

```
func main()  {
    log.Printf("pid:%d\n", os.Getpid())
    command := "/bin/sh"
    if err := syscall.Exec(command, []string{command}, os.Environ()); err != nil {
    }   
}
//output 可以看到PID是一样的，也就是/bin/sh完全替代了go进程的进程信息。这也是linux中exec系统函数本来的功能
root@nicktming:~/go/src/github.com/nicktming/mydocker/test/syscall# go run TestExec.go 
2019/03/25 23:53:52 pid:20872
# echo $$
20872
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
