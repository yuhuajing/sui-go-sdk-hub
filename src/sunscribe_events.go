package src

import (
	"context"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func SubEvent() {
	var ctx = context.Background()
	// create a websocket client, connect to the mainnet websocket endpoint
	var subcli = sui.NewSuiWebsocketClient(constant.WssSuiTestnetEndpoint)

	// receiveMsgCh is a channel to receive Sui event
	receiveMsgCh := make(chan models.SuiEventResponse, 10)

	// SubscribeEvent implements the method `suix_subscribeEvent`, subscribe to a stream of Sui event.
	err := subcli.SubscribeEvent(ctx, models.SuiXSubscribeEventsRequest{
		SuiEventFilter: map[string]interface{}{
			"All": []string{},
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		// receive Sui event
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}

}
