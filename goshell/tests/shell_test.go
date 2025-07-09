package tests

import (
    "os"
    "os/exec"
    "testing"
)

func TestRunBasicCommand(t *testing.T) {
    cmd := exec.Command("ls")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        t.Errorf("Command failed: %v", err)
    }
}
