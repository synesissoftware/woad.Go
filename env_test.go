// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad_test

import (
	woad "github.com/synesissoftware/woad.Go"

	"os"
	"sync"
	"testing"
)

// Variables the colour policy reads. Cleared before each managed case so
// the process environment cannot leak into a row.
var managedVars = []string{
	"CLICOLOR",
	"CLICOLOR_FORCE",
	"FORCE_COLOR",
	"NO_COLOR",
	"TERM",
	"WOAD_COLOR",
	"WOAD_COLOUR",
}

type envPair struct {
	name  string
	value string
}

// Serialises Setenv against the process-wide policy cache. The tests share
// one environment, so they must not overlap.
var envGate sync.Mutex

func captureManaged() map[string]*string {
	saved := make(map[string]*string, len(managedVars))

	for _, name := range managedVars {
		value, present := os.LookupEnv(name)

		if present {
			copy := value
			saved[name] = &copy
		} else {
			saved[name] = nil
		}
	}

	return saved
}

func applySaved(saved map[string]*string) {
	for _, name := range managedVars {
		value := saved[name]

		if value == nil {
			os.Unsetenv(name)
		} else {
			os.Setenv(name, *value)
		}
	}
}

func clearManaged() {
	for _, name := range managedVars {
		os.Unsetenv(name)
	}
}

func assign(
	name string,
	value string,
) {
	os.Setenv(name, value)
}

// Holds envGate across environment edits. Restores the prior environment
// and drops the policy cache on the way out.
func withManaged(
	t *testing.T,
	assignments []envPair,
	body func(),
) {
	t.Helper()

	envGate.Lock()
	defer envGate.Unlock()

	saved := captureManaged()

	defer func() {
		applySaved(saved)
		woad.Refresh()
	}()

	clearManaged()

	for _, item := range assignments {
		assign(item.name, item.value)
	}

	woad.Refresh()

	body()
}
