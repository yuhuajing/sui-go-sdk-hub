package src

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/models"

	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
)

func SendModuleCallObject() {
	var ctx = context.Background()
	if key == "" {
		log.Fatal("Empty key")
	}
	priKey, err := ed25519.PrivateKeyFromSeed(key)
	fmt.Printf("signerAccount.Address: %s\n", AddrFromKey(key))

	gasObj := "0x1b7aaaf599e9d6ff56f257a62d69f58e3fd11430ef2a86b83e4913f36f7354ce"

	rsp, err := cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          AddrFromKey(key),
		PackageObjectId: "0xa84e800e2a8cd8da2bae611750514ae5c3a344104a19088e5ab512785fcea273",
		Module:          "example",
		Function:        "verify",
		TypeArguments:   []interface{}{},
		Arguments: []interface{}{
			"d3e95eec136eb2c9d1ef7ff5ee6f1fd38897362175cd76899c0735a5112081445659834d3259146c03bbb75fa102b8066666466a4c2eef2e015b11969b479415",
			"0475a3b9e2b8b14739f4e66adb08adb20200d5b2c24b53522574b782428b85bb8d10c6ba7134153f3c647b0aae73d5e2a1cb128ce0273e60f2fc030cb4d8dbb810",
			"Hello world!",
			//"0x2cba318afda7b04022cf7522925e11117a966b35771bc9e10679c5b0a0c7739b",
		},
		Gas:       &gasObj,
		GasBudget: "100000000",
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
