package src

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

func GetAllObjectsValues() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	suiObjectResponseQuery := models.SuiObjectResponseQuery{
		// only fetch the effects field
		Options: models.SuiObjectDataOptions{
			ShowType:    true,
			ShowContent: true,
			ShowBcs:     true,
			ShowOwner:   true,
		},
	}
	rsp, err := cli.SuiXGetOwnedObjects(ctx, models.SuiXGetOwnedObjectsRequest{
		Address: "0x2cba318afda7b04022cf7522925e11117a966b35771bc9e10679c5b0a0c7739b",
		Query:   suiObjectResponseQuery,
		Limit:   2,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	AllObjectsPrettyPrint(rsp)
}

type object struct {
	Data []struct {
		Data struct {
			ObjectId string `json:"objectId"`
			Version  string `json:"version"`
			Digest   string `json:"digest"`
			Type     string `json:"type"`
			Owner    struct {
				AddressOwner string `json:"AddressOwner"`
			} `json:"owner"`
			Display struct {
				Data  interface{} `json:"data"`
				Error interface{} `json:"error"`
			} `json:"display"`
			Content struct {
				DataType string `json:"dataType"`
				Type     string `json:"type"`
				Fields   struct {
					Balance string `json:"balance,omitempty"`
					Id      struct {
						Id string `json:"id"`
					} `json:"id"`
					TotalSupply struct {
						Fields struct {
							Value string `json:"value"`
						} `json:"fields"`
						Type string `json:"type"`
					} `json:"total_supply,omitempty"`
				} `json:"fields"`
				HasPublicTransfer bool        `json:"hasPublicTransfer"`
				Disassembled      interface{} `json:"disassembled"`
			} `json:"content"`
			Bcs struct {
				DataType          string `json:"dataType"`
				Type              string `json:"type"`
				HasPublicTransfer bool   `json:"hasPublicTransfer"`
				Version           int    `json:"version"`
				BcsBytes          string `json:"bcsBytes"`
			} `json:"bcs"`
		} `json:"data"`
	} `json:"data"`
	NextCursor  string `json:"nextCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

func AllObjectsPrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}
	var res object
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}

	fmt.Println(res.NextCursor)

}
