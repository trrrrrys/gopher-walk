package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

const Gopher = "ʕ◔ ϖ◔ ʔ"

func main() {
	max, _, err := terminal.GetSize(syscall.Stdin)
	if err != nil {
		panic(err)
	}
	max = max - len([]rune(Gopher))
	i := 0
	left := ""
	right := strings.Repeat(" ", max)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-sig:
			fmt.Fprintf(os.Stdout, "\x1b[1G%s < bye", Gopher)
			return
		default:
			switch (i / max) % 2 {
			case 0:
				left += " "
				right = right[1:]
			case 1:
				left = left[1:]
				right += " "
			}
			fmt.Fprintf(os.Stdout, "\x1b[1G%s%s%s", left, Gopher, right)
			time.Sleep(50 * time.Millisecond)
			i += 1
		}
	}
}
