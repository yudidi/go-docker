# 理解这个运行过程

* 涉及3个进程和一次进程替换

进程1. go-docker进程本身,执行run命令

进程2. 通过/proc/self/exe启动又一个具有隔离ns的功能go-docker进程,借助exec()来执行init命令 //自己启动一个自己然后启动用户提供的进程并替换掉第2个自己

进程3. 用户本来要在容器内启动的进程

进程3是通过完全替换进程2(堆栈,PID等)而得来的，所以最终只剩下进程1和进程3.


* init命令会做什么

1. 作为进程2挂载/proc

2. 用进程3对进程2进行完全替换

# 测试通过程序启动init命令，但是init只打印数据，不启动用户进程。

* 感受下在进程中用"exec.Command("/proc/self/exe", "init").Start()"启动一个进程(这个进程恰好是自己)

运行结果: 进程2不去启动进程3，只是打印一些东西。

```
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
		//return container.RunContainerInitProcess(cmd, nil) // 注意这里并不会去启动用户传入的进程。// init进程是通过
	},
}
```
