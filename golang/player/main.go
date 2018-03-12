package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/btjoker/Spiral-submerged/golang/player/library"
	"github.com/btjoker/Spiral-submerged/golang/player/mp"
)

var lib *library.MusicManager
var id = 1

// 未实现
// var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Genre, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 7 {
			id++
			newMusic := &library.MusicEntry{
				ID:     strconv.Itoa(id),
				Name:   tokens[2],
				Artist: tokens[3],
				Genre:  tokens[4],
				Source: tokens[5],
				Type:   tokens[6],
			}
			lib.Add(newMusic)
		} else {
			fmt.Println("Usage: lib Add <name><artist><genre><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("Usage: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The Music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><genre><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
		q -- exit
	`)
	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if len(tokens) != 2 {
			fmt.Println("Missing parameter")
			continue
		}

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
