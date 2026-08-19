package main

import (
	woad "github.com/synesissoftware/woad.Go"

	"fmt"
)

func main() {
	fmt.Println("the colour is " + woad.FG_GREEN + "green" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_RED + "red" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_YELLOW + "yellow" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_BLUE + "blue" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_MAGENTA + "magenta" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_CYAN + "cyan" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_WHITE + "white" + woad.RESET + ".")
	fmt.Println("the colour is " + woad.FG_BRIGHT_BLACK + "bright black" + woad.RESET + ".")
}
