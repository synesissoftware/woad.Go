// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad

import (
	"io"
	"os"

	"golang.org/x/term"
)

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

// Whether writer is an *os.File whose descriptor is a terminal.
//
// Every other writer, including a nil writer and a nil *os.File, is not
// attached. This is the host check only. It does not apply the colour
// policy, and it does not call SetConsoleMode.
func attached(writer io.Writer) bool {
	if file, isFile := writer.(*os.File); !isFile || file == nil {
		return false
	} else {
		return term.IsTerminal(int(file.Fd()))
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// Whether colour should be emitted for writer.
//
// PolicyForce is yes for every writer. PolicyInhibit is no for every
// writer. PolicyAuto is yes only when writer is a terminal and TERM is
// not dumb (any capitalisation).
//
// On Unix a terminal is ANSI-capable. Windows virtual-terminal gating and
// mintty recognition are not implemented in this alpha; a Windows terminal
// is treated with the Unix rule.
//
// Parameters:
//   - writer — the destination. Attachment is true only for an *os.File whose descriptor is a terminal;
func SupportsColour(writer io.Writer) bool {
	if policy, termIsDumb := colourState(); policy == PolicyForce {
		return true
	} else if policy == PolicyInhibit {
		return false
	} else {
		return attached(writer) && !termIsDumb
	}
}

// The sequence when writer supports colour; otherwise "".
//
// A nil sequence yields "".
//
// Parameters:
//   - writer — the destination. Attachment is true only for an *os.File whose descriptor is a terminal;
//   - sequence — the SGR bytes to emit, or nil;
func SeqFor(
	writer io.Writer,
	sequence *string,
) string {
	if sequence == nil || !SupportsColour(writer) {
		return ""
	} else {
		return *sequence
	}
}

/* ///////////////////////////// end of file //////////////////////////// */
