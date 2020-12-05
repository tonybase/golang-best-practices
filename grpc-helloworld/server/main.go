package main

import (
	"fmt"
	"go/build"
)

func main() {
	ctx := build.Default
	fmt.Printf("%+v", ctx)
}
