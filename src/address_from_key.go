package src

import (
	"encoding/hex"
	osui "github.com/okx/go-wallet-sdk/coins/sui"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
	"log"
)

func AddrFromKey(key string) string {
	p, err := ed25519.PublicKeyFromSeed(key)
	if err != nil {
		log.Printf("Error export pub key: %v", err)
		return ""
	}
	pub := hex.EncodeToString(p)
	//fmt.Println("pub:", pub)
	addr, err := osui.GetAddressByPubKey(pub)
	if err != nil {
		log.Printf("Error export key pair: %v", err)
		return ""
	}
	return addr
}
