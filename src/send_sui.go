package src

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
)

func SendSuiObject() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	var key = ""
	priKey, err := ed25519.PrivateKeyFromSeed(key)
	fmt.Printf("signerAccount.Address: %s\n", AddrFromKey(key))

	rsp, err := cli.TransferSui(ctx, models.TransferSuiRequest{
		Signer:      AddrFromKey(key),
		SuiObjectId: "0xd06dcde26d48f533a093df3fe9eb0fb8ffb6ed357ddec2c027a631052d68029b",
		GasBudget:   "100000000",
		Recipient:   "0x4552b1b58301a39812cac56b20afee17995fc34a2b0ed25ad8da6af2769b23ee",
		Amount:      "1000000000",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// see the successful transaction url: https://explorer.sui.io/txblock/C7iYsH4tU5RdY1KBeNax4mCBn3XLZ5UswsuDpKrVkcH6?network=testnet
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
