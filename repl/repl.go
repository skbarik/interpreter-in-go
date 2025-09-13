package repl

import (
	"bufio"
	"fmt"
	"github.com/skbarik/interpreter-in-go/lexer"
	"github.com/skbarik/interpreter-in-go/token"
	"io"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {

		_, err := fmt.Fprintf(out, prompt)
		if err != nil {
			return
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				return
			}
		}
	}
}
