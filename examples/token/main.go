package main

import (
	"contract-template/sdk"
	"strconv"
	"strings"
)

const MaxSupply = 1000000
const Precision = 3
const Symbol = "TOKEN"
const Creator = "hive:vaultec.vsc"

// Get boolean of whether token has been initialized.
func isInit() bool {
	i := sdk.StateGetObject("isInit")
	return i != nil
}

// Abort execution if token has not been initialized.
func assertInit() {
	if !isInit() {
		// env.Abort()
	}
}

// Get contract owner address and boolean of whether caller is an owner.
func getOwner() (sdk.Address, bool) {
	i := sdk.StateGetObject("owner")
	e := sdk.GetEnv()
	return sdk.Address(*i), *i == e.Caller.Address.String()
}

// Perform a + b addition. Aborts execution if an overflow is detected.
func safeAdd(a, b uint64) uint64 {
	sum := a + b
	// Overflow occurs if sum < a (or equivalently sum < b)
	overflow := sum < a
	if overflow {
		// env.Abort()
	}
	return sum
}

// Perform a - b subtraction. Aborts execution if an overflow is detected.
func safeSub(a, b uint64) uint64 {
	sum := a - b
	// Overflow occurs if sum > a (or equivalently sum < b)
	overflow := sum > a
	if overflow {
		// env.Abort()
	}
	return sum
}

// Increment token balance of an address.
func incBalance(account sdk.Address, amount uint64) {
	oldBal := getBalance(account)
	newBal := safeAdd(oldBal, amount)
	sdk.StateSetObject("accs/"+account.String()+"/bal", strconv.FormatUint(newBal, 10))
}

// Decrement token balance of an address. Aborts execution if insufficient balance.
func decBalance(account sdk.Address, amount uint64) {
	oldBal := getBalance(account)
	if oldBal >= amount {
		newBal := safeSub(oldBal, amount)
		sdk.StateSetObject("accs/"+account.String()+"/bal", strconv.FormatUint(newBal, 10))
	} else {
		// env.Abort("Insufficient balance")
	}
}

// Retrieve token balance of an address.
func getBalance(account sdk.Address) uint64 {
	bal := sdk.StateGetObject("accs/" + account.String() + "/bal")
	if bal == nil {
		return 0
	} else {
		amt, _ := strconv.ParseUint(*bal, 10, 64)
		return amt
	}
}

// Initialize the token contract. Must be called by the Creator address.
//
//go:wasmexport init
func Init(a *string) *string {
	isInit := sdk.StateGetObject("isInit")
	if isInit == nil {
		// env.Abort()
	}
	env := sdk.GetEnv()
	if env.Caller.Address.String() != Creator {
		// env.Abort()
	}
	sdk.StateSetObject("isInit", "1")
	sdk.StateSetObject("supply", "0")
	sdk.StateSetObject("owner", Creator)
	return nil
}

// Mint new tokens to account owner. The caller must be the owner of the token contract.
//
//go:wasmexport mint
func Mint(a *string) *string {
	owner, isOwner := getOwner()
	if !isInit() || !isOwner {
		// env.Abort("Must be owner")
	}
	toMint, err := strconv.ParseUint(*a, 10, 64)
	if err != nil {
		// env.abort("Invalid amount")
	}
	supplyStr := sdk.StateGetObject("supply")
	supply, _ := strconv.ParseUint(*supplyStr, 10, 64)
	newSupply := safeAdd(toMint, supply)
	if newSupply <= MaxSupply {
		sdk.StateSetObject("supply", strconv.FormatUint(newSupply, 10))
		incBalance(owner, toMint)
	} else {
		// env.Abort()
	}
	return nil
}

// Burn tokens from contract caller reducing its current total supply.
//
//go:wasmexport burn
func Burn(a *string) *string {
	assertInit()
	toBurn, err := strconv.ParseUint(*a, 10, 64)
	if err != nil {
		// env.Abort("Invalid amount")
	}
	env := sdk.GetEnv()
	decBalance(env.Caller.Address, toBurn)
	supplyStr := sdk.StateGetObject("supply")
	supply, _ := strconv.ParseUint(*supplyStr, 10, 64)
	newSupply := safeSub(supply, toBurn)
	sdk.StateSetObject("supply", strconv.FormatUint(newSupply, 10))
	return nil
}

// Transfer tokens from caller address to another address. Argument is a comma-separated string of destination address and amount.
//
//go:wasmexport transfer
func Transfer(a *string) *string {
	assertInit()
	params := strings.Split(*a, ",")
	if len(params) < 2 {
		// env.Abort("Invalid number of parameters")
	}
	env := sdk.GetEnv()
	from := env.Caller.Address.String()
	to := params[0]
	amt, err := strconv.ParseUint(params[1], 10, 64)
	if err != nil {
		// env.Abort("Invalid amount")
	}
	decBalance(sdk.Address(from), amt)
	incBalance(sdk.Address(to), amt)
	return nil
}

// Transfer token contract ownership to another address. The caller must be the current owner of the token contract.
//
//go:wasmexport changeOwner
func ChangeOwner(a *string) *string {
	assertInit()
	_, isOwner := getOwner()
	if isOwner {
		sdk.StateSetObject("owner", *a)
	} else {
		// env.Abort()
	}
	return nil
}
