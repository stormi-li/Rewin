package rewin

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func CreateRedisNode(port int, path string) {
	os.MkdirAll(path, 0755)
	//删除配置文件
	truncateFile(path + "/redis.conf")
	//创建配置文件
	appendToFile(path+"/redis.conf", "bind 0.0.0.0\n")
	appendToFile(path+"/redis.conf", "protected-mode no\n")
	appendToFile(path+"/redis.conf", "databases 1\n")
	appendToFile(path+"/redis.conf", "daemonize yes\n")
	appendToFile(path+"/redis.conf", "port "+strconv.Itoa(port)+"\n")
	appendToFile(path+"/redis.conf", "dir "+path+"\n")
	appendToFile(path+"/redis.conf", "always-show-logo yes\n")
	appendToFile(path+"/redis.conf", "loglevel verbose\n")
	appendToFile(path+"/redis.conf", "save 900 1\n")
	appendToFile(path+"/redis.conf", "save 300 10\n")
	appendToFile(path+"/redis.conf", "save 60 10000\n")
	//启动redis
	go func() { exec.Command("cmd", "/C", "start redis-server "+path+"/redis.conf").CombinedOutput() }()
	time.Sleep(100 * time.Millisecond)
}

func truncateFile(filename string) {
	createFileNX(filename)
	file, _ := os.OpenFile(filename, os.O_RDWR, 0644)

	defer file.Close()

	file.Truncate(0)
}

func createFileNX(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func appendToFile(filename string, s string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	s = s + "\n"
	if s == "" {
		return
	}
	_, err = io.WriteString(file, s)
	if err != nil {
		return
	}
}
