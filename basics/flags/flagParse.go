package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "name to greet")

func main() {
    flag.Parse()
    fmt.Printf("Hello, %s!\n", *name)
}
