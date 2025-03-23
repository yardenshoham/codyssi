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
	tracks := make([][]byte, 0, 100)
	for _, track := range bytes.Fields(parts[0]) {
		tracks = append(tracks, track)
	}
	for _, instruction := range bytes.Fields(parts[1]) {
		var a, b int
		_, err := fmt.Sscanf(string(instruction), "%d-%d", &a, &b)
		if err != nil {
			panic(err)
		}
		a--
		b--
		size := min(max(a, b)-min(a, b), len(tracks)-max(a, b))
		for size > 0 {
			tracks[a], tracks[b] = tracks[b], tracks[a]
			a++
			b++
			size--
		}
	}
	testTrack, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		panic(err)
	}
	testTrack--
	fmt.Println(string(tracks[testTrack]))
}
