# init命令
2个功能
1. 在容器中启动用户传入的进程，用户传入的进程会成为ns中第一个启动的进程。
2. 重新挂载proc,使得proc只显示容器的统计和运行信息？TODO

容器内，其实就是指某个ns内的东西。

# exec介绍
linux一共有2种启动新进程的底层调用。
fork:这是启动一个新的进程。这个命令执行完后，会存放父子两个进程。

exec:这是启动一个外部进程，同时这个子进程会完全替换掉启动他的父进程，就好像父进程的唯一意义就是为了启动他这个进程。
这个命令执行完后，就只剩下子进程了，并且这个子进程的PID也是他的父进程的。

 A point worth noting here is that with a call to any of the exec family of functions, the current process image is replaced by a new process image.
 
 > [Linux Processes – Process IDs, fork, execv, wait, waitpid C Functions](https://www.thegeekstuff.com/2012/03/c-process-control-functions/)