package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  strs := strings.Split(str, " ")
  fmt.Println(len(strs))
}