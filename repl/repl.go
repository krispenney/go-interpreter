package repl

import (
  "bufio"
  "fmt"
  "io"
  "interpreter/lexer"
  "interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()

    if (line == "quit") {
      fmt.Printf("Goodbye!\n")
      return
    }

    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      fmt.Printf("%+v\n", tok)
    }
  }
}
