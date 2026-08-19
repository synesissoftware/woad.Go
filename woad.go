// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 20th August 2026
 * Updated: 20th August 2026
 */

// Package woad provides minimal ANSI terminal colour codes for Go.
//
// woad provides the smallest useful set of fixed SGR sequences for library
// authors. It is not a console or TUI framework.
package woad

// Reset all attributes.
const RESET = "\x1b[0m"

// Foreground black.
const FG_BLACK = "\x1b[30m"

// Foreground red.
const FG_RED = "\x1b[31m"

// Foreground green.
const FG_GREEN = "\x1b[32m"

// Foreground yellow.
const FG_YELLOW = "\x1b[33m"

// Foreground blue.
const FG_BLUE = "\x1b[34m"

// Foreground magenta.
const FG_MAGENTA = "\x1b[35m"

// Foreground cyan.
const FG_CYAN = "\x1b[36m"

// Foreground white.
const FG_WHITE = "\x1b[37m"

// Foreground bright black.
const FG_BRIGHT_BLACK = "\x1b[90m"

// Foreground bright red.
const FG_BRIGHT_RED = "\x1b[91m"

// Foreground bright green.
const FG_BRIGHT_GREEN = "\x1b[92m"

// Foreground bright yellow.
const FG_BRIGHT_YELLOW = "\x1b[93m"

// Foreground bright blue.
const FG_BRIGHT_BLUE = "\x1b[94m"

// Foreground bright magenta.
const FG_BRIGHT_MAGENTA = "\x1b[95m"

// Foreground bright cyan.
const FG_BRIGHT_CYAN = "\x1b[96m"

// Foreground bright white.
const FG_BRIGHT_WHITE = "\x1b[97m"

// Background black.
const BG_BLACK = "\x1b[40m"

// Background red.
const BG_RED = "\x1b[41m"

// Background green.
const BG_GREEN = "\x1b[42m"

// Background yellow.
const BG_YELLOW = "\x1b[43m"

// Background blue.
const BG_BLUE = "\x1b[44m"

// Background magenta.
const BG_MAGENTA = "\x1b[45m"

// Background cyan.
const BG_CYAN = "\x1b[46m"

// Background white.
const BG_WHITE = "\x1b[47m"

// Background bright black.
const BG_BRIGHT_BLACK = "\x1b[100m"

// Background bright red.
const BG_BRIGHT_RED = "\x1b[101m"

// Background bright green.
const BG_BRIGHT_GREEN = "\x1b[102m"

// Background bright yellow.
const BG_BRIGHT_YELLOW = "\x1b[103m"

// Background bright blue.
const BG_BRIGHT_BLUE = "\x1b[104m"

// Background bright magenta.
const BG_BRIGHT_MAGENTA = "\x1b[105m"

// Background bright cyan.
const BG_BRIGHT_CYAN = "\x1b[106m"

// Background bright white.
const BG_BRIGHT_WHITE = "\x1b[107m"

/* ///////////////////////////// end of file //////////////////////////// */
