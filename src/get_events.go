package src

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func GetEvents() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)
	rsp, err := cli.SuiGetEvents(ctx, models.SuiGetEventsRequest{
		Digest: "GtVgKE7rYFqb26ahaNw8QkNcjJnZ4qVWRigQnvALziqw",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)

}

func GetLimitedEvents() {
	var ctx = context.Background()
	//fetch list of events for a specified query criteria.
	rsp2, err := cli.SuiXQueryEvents(ctx, models.SuiXQueryEventsRequest{
		SuiEventFilter: models.EventFilterByMoveEventType{
			MoveEventType: "0xa84e800e2a8cd8da2bae611750514ae5c3a344104a19088e5ab512785fcea273::example::VerifiedEvent",
		},
		Limit:           3,
		DescendingOrder: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp2)

}
