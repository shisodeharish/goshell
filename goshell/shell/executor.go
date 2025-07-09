package shell

import (
	"fmt"
	"goshell/utils"
	"os"
	"os/exec"
	"runtime"
)

func ExecuteCommand(input string) {
	cmdName, args := utils.ParseInput(input)
	if cmdName == "" {
		return
	}

	// Handle cross-platform command mapping
	if runtime.GOOS == "windows" {
		switch cmdName {
		case "ls":
			cmdName = "cmd"
			args = append([]string{"/C", "dir"}, args...)
		case "clear":
			cmdName = "cmd"
			args = []string{"/C", "cls"}
		case "echo":
			cmdName = "cmd"
			args = append([]string{"/C", "echo"}, args...)
		}
	}

	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
