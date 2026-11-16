// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 20th August 2026
 * Updated: 17th November 2026
 */

package woad

import "github.com/synesissoftware/ver2go"

const (
	VersionMajor uint16 = 0
	VersionMinor uint16 = 2
	VersionPatch uint16 = 0
	VersionAB    uint16 = ver2go.Alpha1
)

var (
	version              = ver2go.CombineVersion(VersionMajor, VersionMinor, VersionPatch, VersionAB)
	versionString string = ver2go.CalcVersionString(VersionMajor, VersionMinor, VersionPatch, VersionAB)
)

// This library's version as a packed 64-bit integer, formed by
// ver2go.CombineVersion from VersionMajor, VersionMinor, VersionPatch,
// and VersionAB. A later release has a strictly greater value than an
// earlier one that uses the same packing.
func Version() uint64 {
	return version
}

// This library's version as a human-readable string, formed by
// ver2go.CalcVersionString from VersionMajor, VersionMinor, VersionPatch,
// and VersionAB. A final version has the form "MAJOR.MINOR.PATCH". A
// prerelease appends a hyphen and the label, e.g. "0.2.0-alpha1".
func VersionString() string {
	return versionString
}

/* ///////////////////////////// end of file //////////////////////////// */
