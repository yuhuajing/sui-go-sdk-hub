package src

import (
	"encoding/hex"
	"fmt"
	osui "github.com/okx/go-wallet-sdk/coins/sui"
	"log"
)

func NewAccount() {
	keypair, err := osui.GenerateKey()
	if err != nil {
		log.Printf("Error generating key pair: %v", err)
		return
	}
	addr := osui.NewAddress(hex.EncodeToString(keypair.Seed()))
	fmt.Println(addr)

	pri := hex.EncodeToString(keypair.Seed())
	//pri := "e427ae37025f4d1ade8788a967eb1521e8bdb4baf6567b75cd4bb710749eb899"
	addr2 := AddrFromKey(pri)
	fmt.Println(addr2)
}
