package command

import (
	"j/pkg/utils"
	"os"
	"path/filepath"
)

func Init(args []string) {
	utils.CreateDir(filepath.Join(os.Getenv("HOME"), ".j", "version"))
}
