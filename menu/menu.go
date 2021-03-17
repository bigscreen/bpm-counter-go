package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bigscreen/bpm-counter-go/counter"
)

type cmdType string

const (
	cmdStart       cmdType = "s"
	cmdQuit        cmdType = "q"
	cmdUnavailable cmdType = "oops"
)

func Show() {
	fmt.Println("BPM counter")
	fmt.Println("Select one of the options:")
	fmt.Println("(s) Start")
	fmt.Println("(q) Quit")
	fmt.Print("your option (s/q)? ")
	startCLIInput()
}

func startCLIInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error,", err.Error())
			continue
		}
		if len(cmd) < 1 {
			continue
		}

		switch runCLICommand(cmd) {
		case cmdStart:
			Show()
			return
		case cmdUnavailable:
			fmt.Print("Option unavailable, select option (s/q)? ")
			startCLIInput()
			return
		default:
			os.Exit(0)
		}
	}
}

func runCLICommand(cmd string) cmdType {
	cmd = strings.TrimSuffix(cmd, "\n")
	cmd = strings.TrimSpace(cmd)
	switch cmd {
	case "s", "S":
		counter.Start()
		return cmdStart
	case "q", "Q":
		return cmdQuit
	default:
		return cmdUnavailable
	}
}
