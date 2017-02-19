package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/julien-roland/go-randtext/randtext"
)

const text = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
occaecat cupidatat non proident, sunt in culpa qui officia deserunt
mollit anim id est laborum.`

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintln(os.Stderr, "usage: generator number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "generator: number: %s\n", err)
		os.Exit(2)
	}

	model := randtext.NewMarkovModel(strings.NewReader(text))
	b := bufio.NewWriterSize(os.Stdout, 4096)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		fmt.Fprintln(b, model.Sentence())
	}
	b.Flush()
}
