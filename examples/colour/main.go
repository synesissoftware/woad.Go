package main

import (
	woad "github.com/synesissoftware/woad.Go"

	"fmt"
	"os"
)

func paint(
	writer *os.File,
	sequence string,
) string {
	return woad.SeqFor(writer, &sequence)
}

func main() {
	writer := os.Stdout
	lines := []struct {
		sequence string
		word     string
	}{
		{woad.FG_GREEN, "green"},
		{woad.FG_RED, "red"},
		{woad.FG_YELLOW, "yellow"},
		{woad.FG_BLUE, "blue"},
		{woad.FG_MAGENTA, "magenta"},
		{woad.FG_CYAN, "cyan"},
		{woad.FG_WHITE, "white"},
		{woad.FG_BRIGHT_BLACK, "bright black"},
	}

	for _, line := range lines {
		fmt.Println("the colour is " + paint(writer, line.sequence) + line.word + paint(writer, woad.RESET) + ".")
	}
}
