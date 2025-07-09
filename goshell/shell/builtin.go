package shell

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func HandleBuiltins(input string) bool {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return false
	}

	switch parts[0] {
	case "cd":
		path := "~"
		if len(parts) > 1 {
			path = parts[1]
		}
		if path == "~" {
			usr, _ := user.Current()
			path = usr.HomeDir
		}
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("cd error:", err)
		}
		return true
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("pwd error:", err)
		} else {
			fmt.Println(dir)
		}
		return true
	}
	return false
}
