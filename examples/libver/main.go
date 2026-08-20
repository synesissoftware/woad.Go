package main

import (
	"fmt"

	"github.com/synesissoftware/ver2go"
	woad "github.com/synesissoftware/woad.Go"
)

func main() {
	fmt.Printf("woad v%s\n", woad.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
