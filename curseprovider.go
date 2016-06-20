package main

import (
	"fmt"
	"os"

	bp "github.com/nexustix/boilerplate"
	"github.com/nexustix/nxcurse"
)

func main() {
	//version := "V.0-1-0"

	args := os.Args

	switch bp.StringAtIndex(1, args) {
	case "search":
		nxcurse.GetMinecraftModSearchphrase(bp.StringAtIndex(2, args))
		fmt.Printf("<-> search for %s\n", bp.StringAtIndex(2, args))
	case "deps":
		fmt.Printf("<-> depsearch for %s\n", bp.StringAtIndex(2, args))

	}
	fmt.Printf("it works !\n")
}
