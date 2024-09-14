package command

import (
	"fmt"
	"j/pkg/utils"
	"os"
	"path/filepath"
)

func Ls(args []string) {

	versions := getVersions()
	for _, version := range versions {
		fmt.Println(version)
	}

}

func getVersions() []string {
	versions := make([]string, 0)
	root := filepath.Join(utils.HomeDir(), ".j", "version")

	// 遍历目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查当前路径是否为目录
		if info.IsDir() {
			// 获取父目录路径
			parentDir := filepath.Dir(path)

			// 检查父目录是否是起始目录
			if parentDir == root {
				// 如果是起始目录，则说明是直接子目录
				versions = append(versions, filepath.Base(path))
			}
		}

		// 返回nil以继续遍历
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录时发生错误:", err)
	}
	return versions
}
