// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 16th November 2026
 * Updated: 16th November 2026
 */

package woad_test

import (
	woad "github.com/synesissoftware/woad.Go"

	"github.com/stretchr/testify/require"

	"bytes"
	"os"
	"testing"
)

func Test_SeqFor_nil_sequence_is_empty(t *testing.T) {
	withManaged(t, []envPair{{"WOAD_COLOUR", "1"}}, func() {
		require.Equal(t, "", woad.SeqFor(&bytes.Buffer{}, nil))
	})
}

func Test_SeqFor_empty_sequence_is_empty(t *testing.T) {
	withManaged(t, []envPair{{"WOAD_COLOUR", "1"}}, func() {
		sequence := ""

		require.Equal(t, "", woad.SeqFor(&bytes.Buffer{}, &sequence))
	})
}

func Test_non_file_writer_is_not_a_terminal(t *testing.T) {
	withManaged(t, nil, func() {
		var file *os.File

		sequence := woad.FG_RED

		require.Equal(t, woad.PolicyAuto, woad.ColourPolicy())
		require.False(t, woad.SupportsColour(&bytes.Buffer{}))
		require.False(t, woad.SupportsColour(nil))
		require.False(t, woad.SupportsColour(file))
		require.Equal(t, "", woad.SeqFor(&bytes.Buffer{}, &sequence))
	})
}
