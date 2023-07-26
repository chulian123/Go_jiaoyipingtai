package base

import (
	"fmt"
	"testing"
)

func TestWallet_GetAddress(t *testing.T) {
	wallet, err := NewWallet()
	if err != nil {
		return
	}
	address := wallet.GetAddress()
	fmt.Println(string(address))
	key := wallet.GetPriKey()
	fmt.Println(key)
	err = wallet.ResetPriKey(key)
	if err != nil {
		return
	}
	address = wallet.GetAddress()
	fmt.Println(address)
	fmt.Println(string(address))
}

func TestWallet_GetTestAddress(t *testing.T) {
	wallet, _ := NewWallet()
	address := wallet.GetTestAddress()
	fmt.Println(string(address))
	key := wallet.GetPriKey()
	fmt.Println(key)

}
