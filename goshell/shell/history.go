package shell

import (
    "fmt"
    "github.com/peterh/liner"
    "os"
    "os/user"
    "path/filepath"
)

var line *liner.State
var historyFilePath string

func InitHistory() {
    usr, _ := user.Current()
    historyFilePath = filepath.Join(usr.HomeDir, ".goshell_history")
    line = liner.NewLiner()
    line.SetCtrlCAborts(true)
    line.SetCompleter(func(line string) (c []string) {
        return
    })

    f, err := os.Open(historyFilePath)
    if err == nil {
        line.ReadHistory(f)
        f.Close()
    }
}

func CloseHistory() {
    f, err := os.Create(historyFilePath)
    if err != nil {
        return
    }
    line.WriteHistory(f)
    f.Close()
    line.Close()
}

func ReadLine() (string, error) {
    return line.Prompt("goshell> ")
}

func ShowHistory() {
    f, err := os.Open(historyFilePath)
    if err != nil {
        fmt.Println("No history available")
        return
    }
    defer f.Close()
    buf := make([]byte, 1024)
    for {
        n, err := f.Read(buf)
        if n > 0 {
            fmt.Print(string(buf[:n]))
        }
        if err != nil {
            break
        }
    }
}
