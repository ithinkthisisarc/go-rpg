package main

import (
  "fmt"
  "bufio"
  "os"
)

func read(prompt string) string {
  var scanner = bufio.NewScanner(os.Stdin)
  fmt.Printf("%s", prompt)
  scanner.Scan()
  return scanner.Text()
}
