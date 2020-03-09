# init命令
2个功能
1. 在容器中启动用户传入的进程，用户传入的进程会成为ns中第一个启动的进程。
2. 重新挂载proc,使得proc只显示容器的统计和运行信息？TODO



# exec介绍
linux一共有2种启动新进程的底层调用。
fork:这是启动一个新的进程。这个命令执行完后，会存放父子两个进程。

exec:这是启动一个外部进程，同时这个子进程会完全替换掉启动他的父进程，就好像父进程的唯一意义就是为了启动他这个进程。
这个命令执行完后，就只剩下子进程了，并且这个子进程的PID也是他的父进程的。

 A point worth noting here is that with a call to any of the exec family of functions, the current process image is replaced by a new process image.
 
 > [Linux Processes – Process IDs, fork, execv, wait, waitpid C Functions](https://www.thegeekstuff.com/2012/03/c-process-control-functions/)
 
 # 编译和执行结果如下
 ```
[root@192 go-docker]# go build .
[root@192 go-docker]# ./go-docker run --ti /bin/sh
run PID: 16952
init PID: 1 // 从这里可以看到/bin/sh进程id已经被欺骗为1了。
{"level":"info","msg":"用户期望容器执行的第一个command /bin/sh","time":"2020-03-08T03:32:32-04:00"}
sh-4.2#
 ```
