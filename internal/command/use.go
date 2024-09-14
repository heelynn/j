package command

import (
	"fmt"
	"j/pkg/utils"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Use(args []string) {
	//校验是否有命令
	if len(args) < 2 {
		fmt.Println("must have [version]")
		return
	}
	//校验输入版本
	versions := getVersions()
	hasVersion := false
	for _, v := range versions {
		if v == args[1] {
			hasVersion = true
		}
	}
	if !hasVersion {
		fmt.Println("version not found")
		return
	}

	source := filepath.Join(utils.HomeDir(), ".j", "version", args[1])
	shortcut := filepath.Join(utils.HomeDir(), ".j", "java")

	fmt.Println(shortcut)
	// 检查当前操作系统
	platform := runtime.GOOS

	// 根据操作系统创建快捷方式
	switch platform {
	case "windows":
		_, err := os.Stat(shortcut)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("folder '%s' is not exist。\n", shortcut)
			}
		} else {
			err2 := os.Remove(shortcut)
			if err2 != nil {
				fmt.Println("remove shortcut fail : ", err2)
			}
		}

		cmd := exec.Command("cmd.exe", "/c", "mklink", "/j", shortcut, source)
		// 执行命令
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("switch version fail , shortcut fail")
			if err != nil {
				fmt.Println("无法删除目录 '%s': %v", err)
			}
			return
		}
		//utils.CreateWindowsShortcut(source, shortcut)

	case "darwin", "linux":
		// Unix-like系统创建符号链接
		err2 := os.RemoveAll(shortcut)
		if err2 != nil {
			fmt.Println("remove fail , shortcut fail")
		}
		err := os.Symlink(source, shortcut)
		if err != nil {
			fmt.Println("switch version fail , shortcut fail")
		}
	default:
		fmt.Println("this OS is not supported")
	}

	fmt.Println("use the version " + args[1])

}
