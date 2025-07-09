package shell

import (
    "io"
    "os"
    "os/exec"
    "strings"
    "goshell/utils"
)

func ExecutePipedCommand(input string) {
    commands := strings.Split(input, "|")
    var cmds []*exec.Cmd

    for _, cmdStr := range commands {
        name, args := utils.ParseInput(cmdStr)
        if name == "" {
            continue
        }
        cmds = append(cmds, exec.Command(name, args...))
    }

    if len(cmds) == 0 {
        return
    }

    for i := 0; i < len(cmds)-1; i++ {
        r, w := io.Pipe()
        cmds[i].Stdout = w
        cmds[i+1].Stdin = r
    }

    cmds[len(cmds)-1].Stdout = os.Stdout

    for _, cmd := range cmds {
        cmd.Stderr = os.Stderr
    }

    for _, cmd := range cmds {
        _ = cmd.Start()
    }

    for _, cmd := range cmds {
        _ = cmd.Wait()
    }
}
