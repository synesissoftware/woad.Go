package main

import (
	woad "github.com/synesissoftware/woad.Go"
	"github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("woad v%s\n", ver2go.CalcVersionString(woad.VersionMajor, woad.VersionMinor, woad.VersionPatch, woad.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
