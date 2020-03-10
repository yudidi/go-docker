# Mac 下编译 Linux 和 Windows 64位可执行程序

rm ./go-docker
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

scp ./go-docker root@192.168.244.130:/root/go-docker/go-docker


# 改动
为了在用户进程启动前执行mount proc操作。
需要proc/self/exe启动一个新进程,在proc/self/exe进程中mount proc，并且启动用户进程并替换掉self/exe进程。
使得用户进程获得self/exe的NS和PID等信息。

# 编译和运行结果如下
```
=== shell 1====
[root@192 go-docker]# ./go-docker run --it /bin/sh
sh-4.2# ps
   PID TTY          TIME CMD
     1 pts/2    00:00:00 exe
     4 pts/2    00:00:00 sh
     5 pts/2    00:00:00 ps
     
=== shell 2(宿主机)====
[root@192 ~]# ps
Error, do this: mount -t proc proc /proc
[root@192 ~]# mount |grep proc
mount: 读取 mtab 失败: 没有那个文件或目录  
```

## kill 进程1，init进程仍然会继续打印

