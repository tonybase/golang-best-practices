package main

import (
	"fmt"
)

func main() {
	fmt.Println("version:", Version)
	fmt.Println("branch:", Branch)
	fmt.Println("revision:", Revision)
	fmt.Println("build_date:", BuildDate)
}
