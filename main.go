package main

import (
	"fmt"
	"time"

	"github.com/pkg/term"
)

const (
	asciiCodeSpace      = 32
	asciiCodeQUppercase = 81
	asciiCodeQLowercase = 113

	timeoutDuration       = 15 * time.Second
	timeoutBufferDuration = 5 * time.Millisecond
)

func main() {
	fmt.Println("Welcome to BPM counter")
	fmt.Println("- Press <space> to input the beat sample")
	fmt.Println("- Press <q> to exit")
	startCLI()
}

func startCLI() {
	var timeout time.Time
	setTimeout := true
	count := 0

	for {
		c, err := getInputASCIICode()
		if err != nil {
			fmt.Println("Error occurred,", err.Error())
			return
		}

		switch c {
		case asciiCodeSpace:
			{
				if setTimeout {
					timeout = time.Now().Add(timeoutDuration).Add(timeoutBufferDuration)
					setTimeout = false
				}
				if time.Now().Unix() < timeout.Unix() {
					count++
					fmt.Print(" -0- ")
				} else {
					fmt.Print("\n")
					fmt.Printf("Beat Per 15 seconds = %d\n", count)
					fmt.Printf("Beat Per Minute = %d\n", count*4)
					return
				}
			}
		case asciiCodeQUppercase, asciiCodeQLowercase:
			fmt.Println("Bye bye")
			return
		}
	}
}

func getInputASCIICode() (int, error) {
	t, _ := term.Open("/dev/tty")
	_ = term.RawMode(t)

	b := make([]byte, 3)
	numRead, err := t.Read(b)
	_ = t.Restore()
	_ = t.Close()

	if err != nil {
		return 0, err
	}

	if numRead == 1 {
		return int(b[0]), nil
	}

	return 0, err
}
