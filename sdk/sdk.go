package sdk

import (
	_ "contract-template/runtime"
)

//go:wasmimport sdk console.log
func Log(s *string)
