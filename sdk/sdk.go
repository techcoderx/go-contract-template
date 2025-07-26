package sdk

import (
	_ "contract-template/runtime"
)

//go:wasmimport sdk console.log
func Log(s *string) *string

//go:wasmimport sdk db.setObject
func SetObject(key *string, value *string) *string

//go:wasmimport sdk db.getObject
func GetObject(key *string) *string

//go:wasmimport sdk db.delObject
func DelObject(key *string) *string

//go:wasmimport sdk system.call
func Call(method *string, payload *string) *string

//go:wasmimport sdk system.getEnv
func GetEnv(env *string) *string

//go:wasmimport sdk crypto.sha256
func Sha256(val *string) *string

//go:wasmimport sdk crypto.ripemd160
func Ripemd160(val *string) *string

//go:wasmimport sdk hive.getbalance
func GetBalance(account *string, asset *string) *string

//go:wasmimport sdk hive.draw
func Draw(amount *string, asset *string) *string

//go:wasmimport sdk hive.transfer
func Transfer(to *string, amount *string, asset *string) *string

//go:wasmimport sdk hive.withdraw
func Withdraw(to *string, amount *string, asset *string) *string
