package src

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/utils"
)

func GetCheckpoints() {
	var ctx = context.Background()

	rsp, err := cli.SuiGetCheckpoints(ctx, models.SuiGetCheckpointsRequest{
		Limit:           2,
		DescendingOrder: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)

	// fetch Checkpoint 1628214 and print details.
	//rsp2, err := cli.SuiGetCheckpoint(ctx, models.SuiGetCheckpointRequest{
	//	CheckpointID: "1628214",
	//})
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//utils.PrettyPrint(rsp2)
}
