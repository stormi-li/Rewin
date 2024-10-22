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
	TruncateFile(path + "/redis.conf")
	AppendToFile(path+"/redis.conf", "bind 0.0.0.0\n")
	AppendToFile(path+"/redis.conf", "protected-mode no\n")
	AppendToFile(path+"/redis.conf", "databases 1\n")
	AppendToFile(path+"/redis.conf", "daemonize yes\n")
	AppendToFile(path+"/redis.conf", "port "+strconv.Itoa(port)+"\n")
	AppendToFile(path+"/redis.conf", "dir "+path+"\n")
	AppendToFile(path+"/redis.conf", "always-show-logo yes\n")
	AppendToFile(path+"/redis.conf", "loglevel verbose\n")
	AppendToFile(path+"/redis.conf", "save 900 1\n")
	AppendToFile(path+"/redis.conf", "save 300 10\n")
	AppendToFile(path+"/redis.conf", "save 60 10000\n")
	go func() { exec.Command("cmd", "/C", "start redis-server "+path+"/redis.conf").CombinedOutput() }()
	time.Sleep(100 * time.Millisecond)
}

func TruncateFile(filename string) {
	CreateFileNX(filename)
	file, _ := os.OpenFile(filename, os.O_RDWR, 0644)

	defer file.Close()

	file.Truncate(0)
}

func CreateFileNX(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func AppendToFile(filename string, s string) {
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
