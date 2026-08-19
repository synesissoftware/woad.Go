package main

import (
	woad "github.com/synesissoftware/woad.Go"

	"fmt"
)

func main() {
	fmt.Println(woad.FG_GREEN + "ok" + woad.RESET)
	fmt.Println(woad.FG_YELLOW + "warn" + woad.RESET)
	fmt.Println(woad.FG_RED + "error" + woad.RESET)
}
