package src

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
	"unsafe"
)

func SigByKey(key string) {
	if strings.HasPrefix(key, "0x") {
		key = strings.TrimLeft(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}
	message := "Hello world!"
	messageBytes := UnsafeBytes(message)
	_, sigdata := dosig(messageBytes, ecdsaPrivateKey)
	fmt.Printf("SigData: %s\n", sigdata)
}

// SignatureResponse represents the structure of the signature response.
type SignatureResponse struct {
	Address string `json:"address,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Sig     string `json:"sig,omitempty"`
	Version string `json:"version,omitempty"`
}

func dosig(messageBytes []byte, ecdsaPrivateKey *ecdsa.PrivateKey) (string, string) {
	//messageBytes := UnsafeBytes(message)
	msgHash := crypto.Keccak256Hash(messageBytes)
	//hash := accounts.TextHash(msgHash.Bytes())

	signatureBytes, err := crypto.Sign(msgHash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	signatureBytes[64] += 27

	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	fmt.Println(hexutil.Encode(publicKeyBytes))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := crypto.PubkeyToAddress(*pub)

	// Construct the signature response
	res := SignatureResponse{
		Address: rAddress.String(),
		Msg:     string(messageBytes),
		Sig:     hexutil.Encode(signatureBytes),
		Version: "2"}

	// Marshal the response to JSON with proper formatting
	resBytes, err := json.MarshalIndent(res, " ", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(VerifySig(res.Sig, res.Address, res.Msg))

	return res.Sig, string(resBytes)
}

func UnsafeBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
