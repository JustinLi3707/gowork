package main

import (
	"os"

	"github.com/gowork/network"
)

//测试vscode 的代码提交功能

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
