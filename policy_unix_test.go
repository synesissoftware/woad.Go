// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || linux

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad_test

import (
	"os"
	"sync"
	"testing"

	woad "github.com/synesissoftware/woad.Go"

	"github.com/stretchr/testify/require"
	"golang.org/x/term"
)

type endpoints struct {
	terminal *os.File
	master   *os.File
	detached *os.File
}

var (
	endpointsOnce sync.Once
	endpointsHeld endpoints
	endpointsErr  error
)

func testEndpoints(t *testing.T) endpoints {
	t.Helper()

	endpointsOnce.Do(func() {
		master, slave, err := openPty()
		if err != nil {
			endpointsErr = err

			return
		}

		detached, err := os.Open(os.DevNull)
		if err != nil {
			master.Close()
			slave.Close()
			endpointsErr = err

			return
		}

		endpointsHeld = endpoints{
			terminal: slave,
			master:   master,
			detached: detached,
		}
	})

	require.NoError(t, endpointsErr)

	return endpointsHeld
}

func writer(
	t *testing.T,
	terminal bool,
) *os.File {
	t.Helper()

	held := testEndpoints(t)

	if terminal {
		return held.terminal
	} else {
		return held.detached
	}
}

func assertRow(
	t *testing.T,
	row int,
	assignments []envPair,
	terminal bool,
	policy woad.Policy,
	colour bool,
) {
	t.Helper()

	withManaged(t, assignments, func() {
		attached := writer(t, terminal)
		sequence := woad.FG_RED

		require.Equal(t, terminal, term.IsTerminal(int(attached.Fd())), "row %d attachment", row)
		require.Equal(t, policy, woad.ColourPolicy(), "row %d policy", row)
		require.Equal(t, colour, woad.SupportsColour(attached), "row %d colour", row)

		if colour {
			require.Equal(t, sequence, woad.SeqFor(attached, &sequence), "row %d sequence", row)
		} else {
			require.Equal(t, "", woad.SeqFor(attached, &sequence), "row %d sequence", row)
		}
	})
}

func Test_Row_1_none_terminal(t *testing.T) {
	assertRow(t, 1, nil, true, woad.PolicyAuto, true)
}

func Test_Row_2_none_detached(t *testing.T) {
	assertRow(t, 2, nil, false, woad.PolicyAuto, false)
}

func Test_Row_3_TERM_dumb(t *testing.T) {
	assertRow(t, 3, []envPair{{"TERM", "dumb"}}, true, woad.PolicyAuto, false)
}

func Test_Row_4_TERM_DUMB(t *testing.T) {
	assertRow(t, 4, []envPair{{"TERM", "DUMB"}}, true, woad.PolicyAuto, false)
}

func Test_Row_5_WOAD_COLOUR_1(t *testing.T) {
	assertRow(t, 5, []envPair{{"WOAD_COLOUR", "1"}}, false, woad.PolicyForce, true)
}

func Test_Row_6_WOAD_COLOR_always(t *testing.T) {
	assertRow(t, 6, []envPair{{"WOAD_COLOR", "always"}}, false, woad.PolicyForce, true)
}

func Test_Row_7_WOAD_COLOUR_never(t *testing.T) {
	assertRow(t, 7, []envPair{{"WOAD_COLOUR", "never"}}, true, woad.PolicyInhibit, false)
}

func Test_Row_8_WOAD_COLOUR_auto_terminal(t *testing.T) {
	assertRow(t, 8, []envPair{{"WOAD_COLOUR", "auto"}}, true, woad.PolicyAuto, true)
}

func Test_Row_9_WOAD_COLOUR_auto_detached(t *testing.T) {
	assertRow(t, 9, []envPair{{"WOAD_COLOUR", "auto"}}, false, woad.PolicyAuto, false)
}

func Test_Row_10_WOAD_COLOUR_purple(t *testing.T) {
	assertRow(t, 10, []envPair{{"WOAD_COLOUR", "purple"}}, true, woad.PolicyAuto, true)
}

func Test_Row_11_NO_COLOR_beats_CLICOLOR_FORCE(t *testing.T) {
	assertRow(t, 11, []envPair{{"CLICOLOR_FORCE", "1"}, {"NO_COLOR", "1"}}, true, woad.PolicyInhibit, false)
}

func Test_Row_12_CLICOLOR_FORCE(t *testing.T) {
	assertRow(t, 12, []envPair{{"CLICOLOR_FORCE", "1"}}, false, woad.PolicyForce, true)
}

func Test_Row_13_CLICOLOR_FORCE_0_then_CLICOLOR_0(t *testing.T) {
	assertRow(t, 13, []envPair{{"CLICOLOR", "0"}, {"CLICOLOR_FORCE", "0"}}, true, woad.PolicyInhibit, false)
}

func Test_Row_14_FORCE_COLOR(t *testing.T) {
	assertRow(t, 14, []envPair{{"FORCE_COLOR", "1"}}, false, woad.PolicyForce, true)
}

func Test_Row_15_FORCE_COLOR_false_then_CLICOLOR_0(t *testing.T) {
	assertRow(t, 15, []envPair{{"CLICOLOR", "0"}, {"FORCE_COLOR", "false"}}, true, woad.PolicyInhibit, false)
}

func Test_Row_16_FORCE_COLOR_FALSE(t *testing.T) {
	assertRow(t, 16, []envPair{{"FORCE_COLOR", "FALSE"}}, false, woad.PolicyAuto, false)
}

func Test_Row_17_FORCE_COLOR_0(t *testing.T) {
	assertRow(t, 17, []envPair{{"FORCE_COLOR", "0"}}, true, woad.PolicyAuto, true)
}

func Test_Row_18_CLICOLOR_0(t *testing.T) {
	assertRow(t, 18, []envPair{{"CLICOLOR", "0"}}, true, woad.PolicyInhibit, false)
}

func Test_Row_19_CLICOLOR_1(t *testing.T) {
	assertRow(t, 19, []envPair{{"CLICOLOR", "1"}}, true, woad.PolicyAuto, true)
}

func Test_Row_20_empty_WOAD_COLOUR_falls_through(t *testing.T) {
	assertRow(t, 20, []envPair{{"NO_COLOR", "1"}, {"WOAD_COLOUR", ""}}, true, woad.PolicyInhibit, false)
}

func Test_Row_21_empty_NO_COLOR_is_absent(t *testing.T) {
	assertRow(t, 21, []envPair{{"NO_COLOR", ""}}, true, woad.PolicyAuto, true)
}

func Test_Row_22_Refresh_drops_the_cache(t *testing.T) {
	withManaged(t, nil, func() {
		assign("WOAD_COLOUR", "1")
		woad.Refresh()

		attached := writer(t, true)
		sequence := woad.BOLD

		require.Equal(t, woad.PolicyForce, woad.ColourPolicy(), "row 22 before refresh")
		require.True(t, woad.SupportsColour(attached), "row 22 before refresh")

		os.Unsetenv("WOAD_COLOUR")
		assign("NO_COLOR", "1")
		woad.Refresh()

		require.Equal(t, woad.PolicyInhibit, woad.ColourPolicy(), "row 22 after refresh")
		require.False(t, woad.SupportsColour(attached), "row 22 after refresh")
		require.Equal(t, "", woad.SeqFor(attached, &sequence))
		require.True(t, term.IsTerminal(int(attached.Fd())), "row 22 attachment stays a terminal")
	})
}
