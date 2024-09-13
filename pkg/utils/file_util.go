package utils

import (
	"fmt"
	"os"
)

func CreateDir(folderPath string) {
	// 指定需要检查的文件夹路径

	// 检查文件夹是否存在
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Printf("创建文件夹失败: %s\n", err)
		}
		fmt.Printf("文件夹 '%s' 创建成功。\n", folderPath)
	} else if err != nil {
		// 其他错误
		fmt.Printf("检查文件夹状态失败: %s\n", err)
	} else {
		// 文件夹存在
		fmt.Printf("文件夹 '%s' 已存在。\n", folderPath)
	}

}
