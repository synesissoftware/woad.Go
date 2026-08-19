package woad_test

import (
	"strings"

	"github.com/stretchr/testify/require"
	"github.com/synesissoftware/woad.Go"

	"testing"
)

func allCodes() []string {
	return []string{
		woad.RESET,
		woad.FG_BLACK,
		woad.FG_RED,
		woad.FG_GREEN,
		woad.FG_YELLOW,
		woad.FG_BLUE,
		woad.FG_MAGENTA,
		woad.FG_CYAN,
		woad.FG_WHITE,
		woad.FG_BRIGHT_BLACK,
		woad.FG_BRIGHT_RED,
		woad.FG_BRIGHT_GREEN,
		woad.FG_BRIGHT_YELLOW,
		woad.FG_BRIGHT_BLUE,
		woad.FG_BRIGHT_MAGENTA,
		woad.FG_BRIGHT_CYAN,
		woad.FG_BRIGHT_WHITE,
		woad.BG_BLACK,
		woad.BG_RED,
		woad.BG_GREEN,
		woad.BG_YELLOW,
		woad.BG_BLUE,
		woad.BG_MAGENTA,
		woad.BG_CYAN,
		woad.BG_WHITE,
		woad.BG_BRIGHT_BLACK,
		woad.BG_BRIGHT_RED,
		woad.BG_BRIGHT_GREEN,
		woad.BG_BRIGHT_YELLOW,
		woad.BG_BRIGHT_BLUE,
		woad.BG_BRIGHT_MAGENTA,
		woad.BG_BRIGHT_CYAN,
		woad.BG_BRIGHT_WHITE,
	}
}

func Test_RESET(t *testing.T) {
	require.Equal(t, "\x1b[0m", woad.RESET)
}

func Test_FG_RED(t *testing.T) {
	require.Equal(t, "\x1b[31m", woad.FG_RED)
}

func Test_BG_BLUE(t *testing.T) {
	require.Equal(t, "\x1b[44m", woad.BG_BLUE)
}

func Test_all_codes_are_csi_sgr(t *testing.T) {
	for _, code := range allCodes() {
		require.True(t, strings.HasPrefix(code, "\x1b["))
		require.True(t, strings.HasSuffix(code, "m"))
	}
}

func Test_string_concatenation(t *testing.T) {
	concatenated := woad.FG_GREEN + "ok" + woad.RESET

	require.Contains(t, concatenated, "ok")
	require.True(t, strings.HasPrefix(concatenated, woad.FG_GREEN))
}
