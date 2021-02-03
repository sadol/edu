package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// CLI flag definitions here
var LineNumbers = flag.Bool("l", false, "print line numbers")

// tool definition: GO's version of `cat'; closured edition
func cat_enc(showLines bool) (cat func(*bufio.Reader)) {
	var lineNumber int = 1 // state to store

	cat = func(r *bufio.Reader) {
		for {
			if buffer, err := r.ReadBytes('\n'); err == io.EOF {
				break
			} else {
				if showLines == false { // unnumberred version
					fmt.Fprintf(os.Stdout, "%s", buffer) // print to writer(file)
				} else { // numbered version
					fmt.Fprintf(os.Stdout, "%d: \t%s", lineNumber, buffer) // print to writer(file)
					lineNumber++
				}
			}
		}
	}
	return
}

func main() {
	flag.Parse()                   // parse args
	myCat := cat_enc(*LineNumbers) // choose version of the cat function

	if flag.NArg() == 0 { // number of non-flag args of the CLI command
		myCat(bufio.NewReader(os.Stdin)) // read from stdin in case of lack of filenames
	} else { // there are some filenames supplied
		for i := 0; i < flag.NArg(); i++ {
			// post flag args treat as filenames
			if file, err := os.Open(flag.Arg(i)); err != nil {
				fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s.\n",
					os.Args[0], flag.Arg(i), err.Error())
				continue
			} else {
				defer file.Close()
				myCat(bufio.NewReader(file))
			}
		}
	}
}
