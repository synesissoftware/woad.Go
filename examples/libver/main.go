package main

import (
	ver2go "github.com/synesissoftware/ver2go"
	woad "github.com/synesissoftware/woad.Go"

	"fmt"
	"os"
)

// One version line, coloured like the C test.scratch.versions printer:
// library name in blue, version text in green, the Version label in cyan,
// and the packed value in green. Sequences are omitted when writer does
// not support colour.
func versionLine(
	writer *os.File,
	prefix string,
	libname string,
	versionString string,
	packed uint64,
) {
	blue := woad.FG_BLUE
	cyan := woad.FG_CYAN
	green := woad.FG_GREEN
	reset := woad.RESET

	fmt.Fprintf(
		writer,
		"%s%s%s%s: %sv%s%s (%sVersion%s = %s0x%016x%s)\n",
		prefix,
		woad.SeqFor(writer, &blue),
		libname,
		woad.SeqFor(writer, &reset),
		woad.SeqFor(writer, &green),
		versionString,
		woad.SeqFor(writer, &reset),
		woad.SeqFor(writer, &cyan),
		woad.SeqFor(writer, &reset),
		woad.SeqFor(writer, &green),
		packed,
		woad.SeqFor(writer, &reset),
	)
}

func main() {
	writer := os.Stdout

	versionLine(writer, "", "woad.Go", woad.VersionString(), woad.Version())

	fmt.Fprintf(writer, "\nefferent dependencies:\n")

	versionLine(writer, "\t", "ver2go", ver2go.VersionString(), ver2go.Version())
}
