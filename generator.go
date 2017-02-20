package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/julien-roland/go-randtext/randtext"
)

var number = flag.Uint("number", 1, "number of sentences")

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "generator: missing sample text")
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "counter: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	model := randtext.NewMarkovModel(file)
	b := bufio.NewWriterSize(os.Stdout, 4096)
	rand.Seed(time.Now().UnixNano())
	for i := uint(0); i < *number; i++ {
		fmt.Fprintln(b, model.Sentence())
	}
	b.Flush()
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: generator [options] sample")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
