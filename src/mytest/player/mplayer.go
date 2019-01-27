package main

import (
	"bufio"
	"fmt"
	"mytest/player/library"
	"mytest/player/mp"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	fmt.Println("tokens[1]=", tokens[1])
	fmt.Println("tokens[1]==list:", tokens[1] == "list")
	switch tokens[1] {
	case "list":
		fmt.Println("lib.len=", lib.Len())
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		fmt.Println("len(tokens):", len(tokens))
		fmt.Println("tokens:", tokens)
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE:lib add <name><artist><source><type>")
		}
	case "remove":
		fmt.Println("len(tokens):", len(tokens))
		fmt.Println("tokens:", tokens)
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE:lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE:play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist")
		return
	}
	mp.Play(e.Source, e.Type)
}
func main() {
	fmt.Println(`
		Enter following commands to control the player:
			lib list -- View the existing music lib
			lib add <name><artist><source><type> -- Add lib 
			lib remove <name> -- Remove the specif
			play <name> -- Play the specified mus	
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
		fmt.Println("tokens:", tokens)
		fmt.Println("tokens[0]=", tokens[0])
		if tokens[0] == "lib" {
			fmt.Println("1111")
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			fmt.Println("2222")
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}

	}
}
