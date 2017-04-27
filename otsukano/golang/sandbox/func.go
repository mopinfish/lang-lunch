package main

import (
  "os"
  "fmt"
)


func init() {
  var user = os.Getenv("USER")
  if user == "" {
    panic("no value for $USER")
  }
}

func main() {
  init()
  fmt.Println(max(3, 5, 3))
}

// a、b、cの中から最大値を返します。
func max(a, b, c int) int {
  if a > b {
    if (a > c) {
      return a
    }
  }
  if b > c {
    if (b > a) {
      return b
    }
  }
  return c
}

