package command

import (
	"j/pkg/utils"
	"path/filepath"
)

func Init(args []string) {
	utils.CreateDir(filepath.Join(utils.HomeDir(), ".j", "version"))
}
