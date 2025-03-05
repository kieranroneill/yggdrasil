package main

import (
  "fmt"
  "log"
  "os"
)

func main() {
  log.Println(fmt.Sprintf("hello humie! i am version %s", os.Getenv("VERSION")))
}
