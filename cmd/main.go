package main

import (
	"j/internal/args"
	"os"
)

func main() {
	a := os.Args[1:]
	args.SwitchArgs(a)
	// 获取当前用户的主目录
}
