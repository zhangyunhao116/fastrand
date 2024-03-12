//go:build go1.22
// +build go1.22

package fastrand

import (
	_ "unsafe"
)

//go:linkname runtimefastrand runtime.cheaprand
func runtimefastrand() uint32

//go:linkname runtimefastrand64 runtime.cheaprand64
func runtimefastrand64() uint64

//go:linkname runtimefastrandu runtime.cheaprandu
func runtimefastrandu() uint
