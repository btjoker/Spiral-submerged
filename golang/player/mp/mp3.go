package mp

import (
	"fmt"
	"time"
)

// MP3Player .
type MP3Player struct {
	stat     int
	progress int
}

// Play .
func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 musice", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
