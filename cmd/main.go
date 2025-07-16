package main

import (
	"bufio"
	"fmt"
	"github.com/kche0169/KcheFS/fs"
	"os"
	"strings"
)

func main() {
	filesystem := fs.InitFileSystem() // 假设有初始化函数
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		args := strings.Fields(input)
		cmd := args[0]
		switch cmd {
		case "ls":
			filesystem.ListDir(args[1:])
		case "cd":
			filesystem.ChangeDir(args[1:])
		case "mkdir":
			filesystem.CreateDir(args[1:])
		case "touch":
			filesystem.CreateFile(args[1:])
		case "rm":
			filesystem.Remove(args[1:])
		case "exit":
			return
		default:
			fmt.Println("未知命令")
		}
	}
}
