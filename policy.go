// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad

import (
	"os"
	"strings"
	"sync"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

// Prevailing colour policy. The values are the woad 0.3.0 discriminants.
type Policy int

const (
	// Colour depends on the writer. TERM=dumb suppresses colour on this
	// path only.
	PolicyAuto Policy = 0

	// Colour for every writer, including a non-terminal and TERM=dumb.
	PolicyForce Policy = 1

	// Colour for no writer, including a terminal.
	PolicyInhibit Policy = 2
)

/* /////////////////////////////////////////////////////////////////////////
 * constants
 */

var (
	falsey = []string{
		"0",
		"false",
		"never",
		"no",
		"none",
		"off",
	}
	truey = []string{
		"1",
		"always",
		"force",
		"on",
		"true",
		"yes",
	}
)

/* /////////////////////////////////////////////////////////////////////////
 * globals
 */

type environment struct {
	determined bool
	policy     Policy
	termIsDumb bool
}

// Process-wide colour-policy cache. determined starts false. refresh
// clears it and does not itself re-read the environment.
var (
	cacheGate sync.Mutex
	cache     environment
)

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

// The cached policy and the cached TERM=dumb flag.
func colourState() (Policy, bool) {
	cacheGate.Lock()
	defer cacheGate.Unlock()

	ensure()

	return cache.policy, cache.termIsDumb
}

// Fills the cache when it has been dropped. Caller holds cacheGate.
func ensure() {
	if !cache.determined {
		cache.policy = policyFromEnvironment()
		cache.termIsDumb = termIsDumbFromEnvironment()
		cache.determined = true
	}
}

// Whether value equals one of tokens, ignoring case.
func listed(
	value string,
	tokens []string,
) bool {
	for _, token := range tokens {
		if strings.EqualFold(value, token) {
			return true
		}
	}

	return false
}

// The environment value when the variable is set and not empty.
func nonempty(name string) (string, bool) {
	value, present := os.LookupEnv(name)

	if !present || value == "" {
		return "", false
	} else {
		return value, true
	}
}

// The colour policy from the current environment, ignoring the cache.
// The first match wins. Values are not trimmed.
func policyFromEnvironment() Policy {
	if value, present := woadVariable(); present {
		if listed(value, truey) {
			return PolicyForce
		} else if listed(value, falsey) {
			return PolicyInhibit
		} else {
			return PolicyAuto
		}
	}

	if _, present := nonempty("NO_COLOR"); present {
		return PolicyInhibit
	}

	if value, present := nonempty("CLICOLOR_FORCE"); present {
		if value != "0" {
			return PolicyForce
		}
	}

	if value, present := nonempty("FORCE_COLOR"); present {
		if value != "0" && !strings.EqualFold(value, "false") {
			return PolicyForce
		}
	}

	if value, present := nonempty("CLICOLOR"); present {
		if value == "0" {
			return PolicyInhibit
		}
	}

	return PolicyAuto
}

// Whether TERM equals dumb, ignoring case. Absent or empty is not dumb.
func termIsDumbFromEnvironment() bool {
	value, present := nonempty("TERM")

	if present {
		return strings.EqualFold(value, "dumb")
	} else {
		return false
	}
}

// Non-empty WOAD_COLOUR, otherwise non-empty WOAD_COLOR.
func woadVariable() (string, bool) {
	if value, present := nonempty("WOAD_COLOUR"); present {
		return value, true
	} else {
		return nonempty("WOAD_COLOR")
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// The cached colour policy: 0 auto, 1 force, 2 inhibit.
//
// An unset or empty variable is absent. Values are not trimmed. The first
// match wins:
//
//   - non-empty WOAD_COLOUR, else non-empty WOAD_COLOR — truey forces,
//     falsey inhibits, and anything else (including auto) selects auto and
//     does not fall through;
//   - non-empty NO_COLOR inhibits, whatever its text is;
//   - non-empty CLICOLOR_FORCE other than the exact bytes 0 forces;
//   - non-empty FORCE_COLOR other than exact 0 or false (any
//     capitalisation) forces;
//   - CLICOLOR equal to the exact bytes 0 inhibits;
//   - otherwise auto.
//
// Truey: 1, always, force, on, true, yes. Falsey: 0, false, never, no,
// none, off. Both lists are case-insensitive. CLICOLOR and CLICOLOR_FORCE
// compare 0 case-sensitively.
func ColourPolicy() Policy {
	policy, _ := colourState()

	return policy
}

// Drops the cached policy and the cached TERM=dumb flag.
//
// The next query reads the environment again. On Unix there is no console
// cache to drop.
func Refresh() {
	cacheGate.Lock()
	cache.determined = false
	cacheGate.Unlock()
}

/* ///////////////////////////// end of file //////////////////////////// */
