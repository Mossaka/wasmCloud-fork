// Code generated by wit-bindgen-go. DO NOT EDIT.

package wallclock

// This file contains wasmimport and wasmexport declarations for "wasi:clocks@0.2.0".

//go:wasmimport wasi:clocks/wall-clock@0.2.0 now
//go:noescape
func wasmimport_Now(result *DateTime)

//go:wasmimport wasi:clocks/wall-clock@0.2.0 resolution
//go:noescape
func wasmimport_Resolution(result *DateTime)