package shell

import (
	"fmt"
	"os"
	"os/signal"
)

func SetupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println()
			fmt.Print("goshell> ")
		}
	}()
}
