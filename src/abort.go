package main

import "contract-template/env"

//go:wasmexport abortMe
func AbortMe(a *string) *string {
	fname := "abort.go"
	env.Abort(a, &fname, 1, 1)
	ret := "Task failed successfully"
	return &ret
}
