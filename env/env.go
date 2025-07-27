package env

import (
	_ "contract-template/runtime"
)

//go:wasmimport env abort
func Abort(msg *string, file *string, line int32, column int32)
