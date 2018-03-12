package mp

import (
	"fmt"
	"time"
)

// WAVPlayer .
type WAVPlayer struct {
	stat     int
	progress int
}

// Play .
func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV musice", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
