package main

import (
	"fmt"

	"github.com/tonybase/golang-best-practices/build-version/version"
)

func main() {
	fmt.Println("version:", version.Version)
	fmt.Println("branch:", version.Branch)
	fmt.Println("revision:", version.Revision)
	fmt.Println("build_date:", version.BuildDate)
}
