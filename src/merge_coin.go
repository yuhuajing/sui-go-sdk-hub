package src

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
)

func SendMergeCoinObject() {
	var ctx = context.Background()

	priKey, err := ed25519.PrivateKeyFromSeed(key)
	fmt.Printf("signerAccount.Address: %s\n", AddrFromKey(key))

	gasObj := "0xd06dcde26d48f533a093df3fe9eb0fb8ffb6ed357ddec2c027a631052d68029b"

	rsp, err := cli.MergeCoins(ctx, models.MergeCoinsRequest{
		Signer:      AddrFromKey(key),
		PrimaryCoin: "0x1b7aaaf599e9d6ff56f257a62d69f58e3fd11430ef2a86b83e4913f36f7354ce",
		CoinToMerge: "0xbcca4db34f6b7dffe4defc184287123b43577e4318a6bea13c8ebbd045a58b20",
		Gas:         &gasObj,
		GasBudget:   "100000000",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// see the successful transaction url: https://explorer.sui.io/txblock/CD5hFB4bWFThhb6FtvKq3xAxRri72vsYLJAVd7p9t2sR?network=testnet
	rsp2, err := cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      priKey,
		// only fetch the effects field
		Options: models.SuiTransactionBlockOptions{
			ShowInput:    true,
			ShowRawInput: true,
			ShowEffects:  true,
		},
		RequestType: "WaitForLocalExecution",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp2)
}
