# Mac 下编译 Linux 和 Windows 64位可执行程序

rm ./go-docker
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

scp ./go-docker root@192.168.244.130:/root/go-docker/go-docker


# 改动
用户进程启动前执行mount proc会导致宿主机挂载的proc失效。
我们需要避免这种影响。

> [namespace里面mount /proc 后,退出后要重新mount](https://blog.csdn.net/qq_27068845/article/details/90705925)

# 编译和运行结果如下
```
=== shell 1====
[root@192 go-docker]# ./go-docker run --it /bin/sh
sh-4.2# ps
   PID TTY          TIME CMD
     1 pts/0    00:00:00 exe
     4 pts/0    00:00:00 sh
     5 pts/0    00:00:00 ps
     
=== shell 2(宿主机)==== // 已经没有影响了
[root@192 ~]# mount |grep proc
proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
systemd-1 on /proc/sys/fs/binfmt_misc type autofs (rw,relatime,fd=31,pgrp=1,timeout=0,minproto=5,maxproto=5,direct,pipe_ino=13811)
[root@192 ~]# ps
   PID TTY          TIME CMD
  1340 pts/1    00:00:00 bash
  1371 pts/1    00:00:00 ps
```

## kill 进程1，init进程仍然会继续打印

