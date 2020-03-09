# 理解这个运行过程

* 涉及3个进程和一次进程替换

进程1. go-docker进程本身,执行run命令

进程2(init进程). 通过/proc/self/exe启动又一个具有隔离ns的功能go-docker进程,借助exec()来执行init命令 //自己启动一个自己然后启动用户提供的进程并替换掉第2个自己

进程3. 用户本来要在容器内启动的进程

进程3是通过完全替换进程2(堆栈,PID等)而得来的，所以最终只剩下进程1和进程3.


* init命令会做什么

1. 作为进程2挂载/proc

2. 用进程3对进程2进行完全替换

# 测试通过程序启动init命令，但是init只打印数据，不启动用户进程。

* 感受下在进程中用"exec.Command("/proc/self/exe", "init").Start()"启动一个进程(`自己调用自己`,但是执行不同的命令)

运行结果: init进程不去启动进程3，只是打印一些东西。

```
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		fmt.Printf("init PID: %d \n", os.Getpid())
		//cmd := context.Args().Get(0)
		for {
			fmt.Println(time.Now())
			time.Sleep(5 * time.Second)
		}
		return nil
		//return container.RunContainerInitProcess(cmd, nil)
	},
}
```
# 编译和运行结果如下
```
[root@192 go-docker]# go build .
[root@192 go-docker]# ./go-docker run --ti /bin/sh
run PID: 17489
init PID: 1
2020-03-08 03:51:25.170607874 -0400 EDT m=+0.000872106
2020-03-08 03:51:30.17159265 -0400 EDT m=+5.001856986
```

## kill 进程1，init进程仍然会继续打印

