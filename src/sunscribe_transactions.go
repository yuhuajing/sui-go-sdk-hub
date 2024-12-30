package src

import (
	"context"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func SubTransactions() {
	var ctx = context.Background()
	// create a websocket client, connect to the mainnet websocket endpoint
	var subcli = sui.NewSuiWebsocketClient(constant.WssSuiTestnetEndpoint)

	// receiveMsgCh is a channel to receive Sui transaction effects
	receiveMsgCh := make(chan models.SuiEffects, 10)

	// SubscribeTransaction implements the method `suix_subscribeTransaction`, subscribe to a stream of Sui transaction effects.
	err := subcli.SubscribeTransaction(ctx, models.SuiXSubscribeTransactionsRequest{
		//TransactionFilter: models.TransactionFilterByFromAddress{
		//	FromAddress: "0x0000000000000000000000000000000000000000000000000000000000000000",
		//},
		TransactionFilter: models.TransactionFilterByToAddress{
			ToAddress: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		// receive Sui transaction effects
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}
}
