package main

import rewin "github.com/stormi-li/Rewin"

func main() {
	rewin.CreateRedisNode(6379,".\\node")
}