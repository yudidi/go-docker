# Mac 下编译 Linux 和 Windows 64位可执行程序

rm ./go-docker
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

scp ./go-docker root@192.168.244.130:/root/go-docker/go-docker


# 功能和代码
直接使用exec.Command启动用户传入的命令，并设置NS.
```
	cmd := exec.Command(command)
	// for kinds of namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
```

# 编译和运行结果如下
可以看到虽然NS已经隔离了，但是ps查看的信息是没有被隔离的，因为父子进程的挂载视图是一致的，且ps等命令的数据源就是/proc这个挂载点挂载的proc类型的设备文件。
```
===shell 1====
[root@localhost go-docker]# ./go-docker run --it /bin/sh
sh-4.2# ps -ef |grep sh
root       1413   1409  0 03:28 pts/0    00:00:00 -bash
root       1440   1413  0 03:28 pts/0    00:00:00 ./go-docker run --it /bin/sh
root       1443   1440  0 03:28 pts/0    00:00:00 /bin/sh
root       1447   1443  0 03:29 pts/0    00:00:00 grep sh
====shell 2(宿主机)====
root       1413   1409  0 03:28 pts/0    00:00:00 -bash
root       1440   1413  0 03:28 pts/0    00:00:00 ./go-docker run --it /bin/sh
root       1443   1440  0 03:28 pts/0    00:00:00 /bin/sh
root       1480   1454  0 03:31 pts/1    00:00:00 grep --color=auto sh
```

