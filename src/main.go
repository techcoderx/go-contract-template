// Proof of Concept VSC Smart Contract in Golang
//
// Build command: tinygo build -o main.wasm -gc=custom -scheduler=none -panic=trap -no-debug -target=wasm-unknown main.go
// Inspect Output: wasmer inspect main.wasm
// Run command (only works w/o SDK imports): wasmedge run main.wasm entrypoint 0
//
// Caveats:
// - Go routines, channels, and defer are disabled
// - panic() always halts the program, since you can't recover in a deferred function call
// - must import sdk or build fails
// - to mark a function as a valid entrypoint, it must be manually exported (//go:wasmexport <entrypoint-name>)
//
// TODO:
// - when panic()ing, call `env.abort()` instead of executing the unreachable WASM instruction
// - Remove _initalize() export & double check not necessary

package main

import (
	"contract-template/sdk"
	"strings"
)

//go:wasmexport entrypoint
func Entrypoint(a *string) *string {
	// panic("test")
	msg := "Entrypoint logged"
	sdk.Log(a)
	sdk.Log(&msg)
	return a
}

//go:wasmexport setString
func SetString(a *string) *string {
	key := "myString"
	sdk.SetObject(&key, a)
	return a
}

//go:wasmexport getString
func GetString(a *string) *string {
	key := "myString"
	return sdk.GetObject(&key)
}

//go:wasmexport clearString
func ClearString(a *string) *string {
	key := "myString"
	return sdk.DelObject(&key)
}

//go:wasmexport dumpEnv
func DumpEnv(a *string) *string {
	varContractId := "contract_id"
	varSender := "msg.sender"
	varReqAuths := "msg.required_auths"
	varReqPostingAuths := "msg.required_posting_auths"
	varAnchrId := "anchor.id"
	varAnchrBlk := "anchor.block"
	varAnchrHeight := "anchor.height"
	varAnchrTs := "anchor.timestamp"
	varAnchrTxIdx := "anchor.tx_index"
	varAnchrOpIdx := "anchor.op_index"
	sdk.Log(sdk.GetEnv(&varContractId))
	sdk.Log(sdk.GetEnv(&varSender))
	sdk.Log(sdk.GetEnv(&varReqAuths))
	sdk.Log(sdk.GetEnv(&varReqPostingAuths))
	sdk.Log(sdk.GetEnv(&varAnchrId))
	sdk.Log(sdk.GetEnv(&varAnchrBlk))
	sdk.Log(sdk.GetEnv(&varAnchrHeight))
	sdk.Log(sdk.GetEnv(&varAnchrTs))
	sdk.Log(sdk.GetEnv(&varAnchrTxIdx))
	sdk.Log(sdk.GetEnv(&varAnchrOpIdx))
	ret := "Dump complete"
	return &ret
}

//go:wasmexport sha256
func Sha256(a *string) *string {
	return sdk.Sha256(a)
}

//go:wasmexport ripemd160
func Ripemd160(a *string) *string {
	return sdk.Ripemd160(a)
}

//go:wasmexport getHiveBalance
func GetHiveBalance(a *string) *string {
	asset := "hive"
	return sdk.GetBalance(a, &asset)
}

//go:wasmexport getHbdBalance
func GetHbdBalance(a *string) *string {
	asset := "hbd"
	return sdk.GetBalance(a, &asset)
}

//go:wasmexport getHiveConsBalance
func GetHiveConsBalance(a *string) *string {
	asset := "hive_consensus"
	return sdk.GetBalance(a, &asset)
}

//go:wasmexport drawHive
func DrawHive(a *string) *string {
	asset := "hive"
	return sdk.Draw(a, &asset)
}

//go:wasmexport drawHbd
func DrawHbd(a *string) *string {
	asset := "hbd"
	return sdk.Draw(a, &asset)
}

//go:wasmexport transferHive
func TransferHive(a *string) *string {
	asset := "hive"
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		err := "invalid payload"
		return &err
	}
	return sdk.Transfer(&params[0], &params[1], &asset)
}

//go:wasmexport transferHbd
func TransferHbd(a *string) *string {
	asset := "hbd"
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		err := "invalid payload"
		return &err
	}
	return sdk.Transfer(&params[0], &params[1], &asset)
}

//go:wasmexport withdrawHive
func WithdrawHive(a *string) *string {
	asset := "hive"
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		err := "invalid payload"
		return &err
	}
	return sdk.Withdraw(&params[0], &params[1], &asset)
}

//go:wasmexport withdrawHbd
func WithdrawHbd(a *string) *string {
	asset := "hbd"
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		err := "invalid payload"
		return &err
	}
	return sdk.Withdraw(&params[0], &params[1], &asset)
}
