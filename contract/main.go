package main

import (
	"encoding/hex"
	"strconv"
	"strings"
	"test-contract/sdk"
)

func main() {

}

//go:wasmexport entrypoint
func Entrypoint(a *string) *string {
	return a
}

//go:wasmexport helloWorld
func HelloWorld(a *string) *string {
	ret := "Hello world"
	return &ret
}

//go:wasmexport setString
func SetString(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Abort("invalid payload")
	}
	sdk.StateSetObject(params[0], params[1])
	return a
}

//go:wasmexport getString
func GetString(a *string) *string {
	return sdk.StateGetObject(*a)
}

//go:wasmexport clearString
func ClearString(a *string) *string {
	sdk.StateDeleteObject(*a)
	return nil
}

//go:wasmexport dumpEnv
func DumpEnv(a *string) *string {
	env := sdk.GetEnvStr()
	sdk.Log("Dump completed")
	return &env
}

//go:wasmexport dumpEnvKey
func DumpEnvKey(a *string) *string {
	return sdk.GetEnvKey(*a)
}

//go:wasmexport getHiveBalance
func GetHiveBalance(a *string) *string {
	ret := strconv.FormatInt(sdk.GetBalance(sdk.Address(*a), sdk.AssetHive), 10)
	return &ret
}

//go:wasmexport getHbdBalance
func GetHbdBalance(a *string) *string {
	ret := strconv.FormatInt(sdk.GetBalance(sdk.Address(*a), sdk.AssetHbd), 10)
	return &ret
}

//go:wasmexport getHiveConsBalance
func GetHiveConsBalance(a *string) *string {
	ret := strconv.FormatInt(sdk.GetBalance(sdk.Address(*a), sdk.AssetHiveCons), 10)
	return &ret
}

//go:wasmexport drawHive
func DrawHive(a *string) *string {
	amt, err := strconv.ParseInt(*a, 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveDraw(amt, sdk.AssetHive)
	return nil
}

//go:wasmexport drawHbd
func DrawHbd(a *string) *string {
	amt, err := strconv.ParseInt(*a, 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveDraw(amt, sdk.AssetHbd)
	return nil
}

//go:wasmexport transferHive
func TransferHive(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Abort("invalid payload")
	}
	amt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveTransfer(sdk.Address(params[0]), amt, sdk.AssetHive)
	return nil
}

//go:wasmexport transferHbd
func TransferHbd(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Abort("invalid payload")
	}
	amt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveTransfer(sdk.Address(params[0]), amt, sdk.AssetHbd)
	return nil
}

//go:wasmexport withdrawHive
func WithdrawHive(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Abort("invalid payload")
	}
	amt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveWithdraw(sdk.Address(params[0]), amt, sdk.AssetHive)
	return nil
}

//go:wasmexport withdrawHbd
func WithdrawHbd(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Abort("invalid payload")
	}
	amt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		sdk.Abort("invalid amount")
	}
	sdk.HiveWithdraw(sdk.Address(params[0]), amt, sdk.AssetHbd)
	return nil
}

//go:wasmexport abortMe
func AbortMe(a *string) *string {
	sdk.Abort(*a)
	ret := "Task failed successfully"
	return &ret
}

//go:wasmexport revertMe
func RevertMe(a *string) *string {
	sdk.Revert(*a, "symbol_here")
	ret := "Task failed successfully"
	return &ret
}

//go:wasmexport contractGetString
func ContractGetString(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Revert("invalid payload", "invalid_payload")
	}
	return sdk.ContractStateGet(params[0], params[1])
}

//go:wasmexport contractCall
func ContractCall(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 3 {
		sdk.Revert("invalid payload", "invalid_payload")
	}
	return sdk.ContractCall(params[0], params[1], params[2], &sdk.ContractCallOptions{
		Intents: []sdk.Intent{{
			Type: "transfer.allow",
			Args: map[string]string{
				"token": "hive",
				"limit": "1.000",
			},
		}},
	})
}

//go:wasmexport infiniteRecursion
func InfRecursion(a *string) *string {
	contractId := sdk.GetEnvKey("contract.id")
	return sdk.ContractCall(*contractId, "infiniteRecursion", "a", nil)
}

//go:wasmexport createKey
func CreateKey(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Revert("invalid payload", "invalid_payload")
	}
	status := sdk.TssCreateKey(params[0], params[1])
	return &status
}

//go:wasmexport getKey
func GetKey(a *string) *string {
	ret := sdk.TssGetKey(*a)
	return &ret
}

//go:wasmexport signKey
func SignKey(a *string) *string {
	params := strings.Split((*a), ",")
	if len(params) < 2 {
		sdk.Revert("invalid payload", "invalid_payload")
	}
	msg, _ := hex.DecodeString(params[1])

	sdk.TssSignKey(params[0], msg)

	ret := ""
	return &ret
}
