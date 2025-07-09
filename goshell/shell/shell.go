package shell

import (
    "fmt"
    "strings"
)

func Start() {
    SetupSignalHandler()
    InitHistory()
    defer CloseHistory()

    for {
        input, err := ReadLine()
        if err != nil {
            fmt.Println()
            break
        }

        input = strings.TrimSpace(input)
        if input == "" {
            continue
        }

        if input == "exit" {
            break
        }

        line.AppendHistory(input)

        if input == "history" {
            ShowHistory()
            continue
        }

        if HandleBuiltins(input) {
            continue
        }

        if strings.Contains(input, "|") {
            ExecutePipedCommand(input)
        } else {
            ExecuteCommand(input)
        }
    }
}
