package woad_test

import (
	"github.com/synesissoftware/woad.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 2
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x4001
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, woad.VersionMajor)
	require.Equal(t, Expected_VersionMinor, woad.VersionMinor)
	require.Equal(t, Expected_VersionPatch, woad.VersionPatch)
	require.Equal(t, Expected_VersionAB, woad.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0002_0000_4001), woad.Version())
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.2.0-alpha1", woad.VersionString())
}
