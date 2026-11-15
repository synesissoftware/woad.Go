// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 20th August 2026
 * Updated: 17th November 2026
 */

// Package woad provides minimal ANSI terminal colour codes for Go.
//
// woad provides the smallest useful set of fixed SGR sequences for library
// authors, and the portable half of the colour policy: cached environment
// inspection, TERM=dumb, and conditional emission for an *os.File.
// It is not a console or TUI framework. Windows virtual-terminal gating is
// a later alpha.
package woad

/* ///////////////////////////// end of file //////////////////////////// */
