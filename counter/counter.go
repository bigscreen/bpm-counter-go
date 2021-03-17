package counter

import (
	"fmt"
	"time"

	"github.com/pkg/term"
)

const (
	asciiCodeSpace      = 32
	asciiCodeQUppercase = 81
	asciiCodeQLowercase = 113

	timeoutInSeconds      = 15
	timeoutDuration       = timeoutInSeconds * time.Second
	timeoutBufferDuration = 5 * time.Millisecond

	beatConstraint = 4
)

func Start() {
	showInstructions()
	startCounter()
}

func showInstructions() {
	fmt.Println("- Press <space> to start and input the beat sample")
	fmt.Println("- Press <q> to back to menu")
}

func startCounter() {
	var timeout time.Time
	setTimeout := true
	beat := int32(0)

	for {
		c, err := getInputASCIICode()
		if err != nil {
			fmt.Println("Error, ", err.Error())
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
					beat++
					showBeatFeedback(beat)
				} else {
					showResults(beat)
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

func showBeatFeedback(b int32) {
	beat := b % beatConstraint
	if beat == 0 {
		beat = beatConstraint
	}
	fmt.Printf(" -%d- ", beat)
}

func showResults(b int32) {
	fmt.Print("\n")
	fmt.Printf("Beat Per 15 seconds = %d\n", b)
	fmt.Printf("Beat Per Minute = %d\n", (60/timeoutInSeconds)*b)
}
