package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func deleteComments(file string) string {

	/* Lines with // could contain special symbols such as .,[] which we are using for transpiling to JS
	   It's the reason we delet them before running transpiling

	 */

	r, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	SLASH := "/"

	slashCode := []byte(SLASH)[0]
	commentCodeSlashes := slashCode*slashCode+slashCode+slashCode

	scanner.Split(bufio.ScanLines)
	var prog string
	var line string

	for scanner.Scan() {
		line = scanner.Text()
		code := line[0]*line[1]+line[0]+line[1]

		if code == commentCodeSlashes {
			continue
		}
		prog += line
	}

    return prog
}

func transpile(prog string, w io.Writer) error {
	
	var (
		fpos uint   = 0                  
		size uint   = 30000
		plen  = uint(len(prog))
		)

	fmt.Fprint(w, "//The script is generated. Do not edit manually\n\n")
	fmt.Fprint(w, "\nlet fs = require('fs');")
	fmt.Fprint(w, "\nfunction getChar() {\nlet buf = Buffer.alloc(1)\n  fs.readSync(0, buf, 0, 1); return buf.toString('ascii').charCodeAt(0);}")
	fmt.Fprint(w, "\nfunction checkedDecrement(d){ if (d === 0) {return 255} return --d} ")
	fmt.Fprint(w, "\nfunction checkedIncrement(d){ if (d === 255) {return 0} return ++d} ")
	fmt.Fprintf(w, "\nlet d=new Array(%d);", size)
	fmt.Fprint(w, "\nlet i=0; let output=\"\";")
	fmt.Fprintf(w, "\nfor(a=0; a<%d; a++){d[a]=0}", size-1)

	for fpos < plen {
		switch prog[fpos] {
		case '+':
			fmt.Fprint(w, "d[i]=checkedIncrement(d[i]);")
		case '-':
			fmt.Fprint(w, "d[i]=checkedDecrement(d[i]);")
		case '>':
		    fmt.Fprint(w, "i++;")
		case '<':
			fmt.Fprint(w, "i--;")
		case '.':
			fmt.Fprint(w, "\n\toutput+=String.fromCharCode(d[i]);")
		case ',':
			fmt.Fprint(w," d[i] = getChar(); console.log('CHAR:',d[i]);")
		case '[':
			fmt.Fprint(w, "\nwhile(d[i]){")
		case ']':
			fmt.Fprint(w, "\n}\n")
		}
		fpos++
	}

	fmt.Fprint(w, "\nconsole.log(output);")
	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%v\n", "Error: programm accepts only 1 argument: filename of bf program")
		os.Exit(-1)
	}

	p := deleteComments(os.Args[1])

	err := transpile(p, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%b\n", err)
		os.Exit(-1)
	}

}


