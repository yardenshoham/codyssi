package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 3 {
		panic("bad input")
	}
	tracks := make([]string, 0, 100)
	for _, track := range bytes.Fields(parts[0]) {
		tracks = append(tracks, string(track))
	}
	instructions := bytes.Fields(parts[1])
	for i, instruction := range instructions {
		var a, b, c int
		_, err := fmt.Sscanf(string(instruction)+" "+string(instructions[(i+1)%len(instructions)]), "%d-%d %d", &a, &b, &c)
		if err != nil {
			panic(err)
		}
		a--
		b--
		c--
		tracks[b], tracks[c], tracks[a] = tracks[a], tracks[b], tracks[c]
	}
	testTrack, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		panic(err)
	}
	testTrack--
	fmt.Println(tracks[testTrack])
}
