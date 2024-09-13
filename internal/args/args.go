package args

import (
	"fmt"
	"j/internal/command"
)

var prompt string = `COMMANDS:
	init			initialize java directory
	ls			show version
	use [version]		use version
`

func SwitchArgs(args []string) {
	if len(args) == 0 {
		fmt.Println(prompt)
		return
	}
	switchFunction(args)

}
func switchFunction(args []string) {
	switch args[0] {
	case "init":
		command.Init(args)
	case "ls":
		command.Ls(args)
	case "use":
		command.Use(args)
	default:
		fmt.Println("Invalid argument\n" + prompt)
	}
}
