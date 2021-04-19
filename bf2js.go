package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/the2pizza/bf2js/stack"
	"io"
	"io/ioutil"
	"os"
)

func eval(r io.Reader, i io.Reader, w io.Writer) error {
	prog, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	input := bufio.NewReader(i) // buffered reader for `,` requests

	var (
		fpos uint   = 0                  // file position
		dpos uint   = 0                  // data position
		size uint   = 30000              // size of data card
		plen uint   = uint(len(prog))    // programme length
		data []byte = make([]byte, size) // data card with `size` items
		promt = 0
	)

	jumps, err := bfJumps(prog)  // pre-computed jumps

	if err != nil {
		return err
	}

	for fpos < plen {
		switch prog[fpos] {
		case '+': // increment at current position
			data[dpos]++
		case '-': // decrement at current position
			data[dpos]--
		case '>': // move to next position
			if dpos == size-1 {
				dpos = 0
			} else {
				dpos++
			}
		case '<': // move to previous position
			if dpos == 0 {
				dpos = size - 1
			} else {
				dpos--
			}
		case '.': // output value of current position
			fmt.Fprintf(w, "%c", data[dpos])
		case ',': // read value into current position
		    if promt == 0 {
				fmt.Fprintf(w, "%v: ", "In")
				promt = 1
			}
			if data[dpos], err = input.ReadByte(); err != nil {
				os.Exit(0)
			}

		case '[': // if current position is false, skip to ]
			if data[dpos] == 0 {
				fpos = jumps[fpos]
			}
		case ']': // if at current position true, return to [
			if data[dpos] != 0 {
				fpos = jumps[fpos]
			}
		}
		fpos++
	}
	return nil
}


func bfJumps(prog []byte) (map[uint]uint, error) {
	var (
		stack *stack.Stack = stack.New()
		jumps map[uint]uint = make(map[uint]uint)

		plen uint = uint(len(prog))
		fpos uint = 0
	)

	for fpos < plen {
		switch prog[fpos] {
		case '[':
			stack.Push(fpos)
		case ']':
			tget, err := stack.Pop()
			if err != nil {
				return nil, errors.New(
					"unexpected closing bracket",
				)
			}
			jumps[tget] = fpos
			jumps[fpos] = tget
		}
		fpos++
	}

	_, err := stack.Pop()
	if err == nil {
		return nil, errors.New(
			"excessive opening brackets",
		)
	}

	fmt.Println("Jumps", jumps)
	return jumps, nil
}



func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%v\n", "Error: programm accepts only 1 argument")
		os.Exit(-1)
	}

	r, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	err = eval(r, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%b\n", err)
		os.Exit(-1)
	}

}


