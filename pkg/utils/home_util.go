package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func HomeDir() string {
	platform := runtime.GOOS

	// 根据操作系统创建快捷方式
	switch platform {
	case "darwin", "linux":
		return filepath.Join(os.Getenv("HOME"))
	default:
		currentUserHomeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("get user home dir fail")
			return ""
		}
		return currentUserHomeDir
	}

}
func CreateWindowsShortcut(source, shortcut string) {
	cmd := exec.Command("cmd.exe", "/c", "powershell -Command \"New-Item -ItemType SymbolicLink -Path '"+shortcut+"' -Name '.lnk' -Target '"+source+"'\"")
	err := cmd.Run()
	if err != nil {
		fmt.Println("create shortcut fail")
	}

}
