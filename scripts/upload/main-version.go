package main

import (
	"fmt"
	packageVersion "properties-cli/version"
)

func main() {
	fmt.Printf("%v\n", packageVersion.Version)
}
