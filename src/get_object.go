package src

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func GetObjectValues() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	rsp, err := cli.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: "0xad5519de766a8577f9b993265c876b0d0cd75d9a43543764b26eac1c8e794f79",
		// only fetch the effects field
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

type Object struct {
	Data struct {
		ObjectId string `json:"objectId"`
		Version  string `json:"version"`
		Digest   string `json:"digest"`
		Type     string `json:"type"`
		Owner    struct {
			AddressOwner string `json:"AddressOwner"`
		} `json:"owner"`
		PreviousTransaction string `json:"previousTransaction"`
		Display             struct {
			Data  interface{} `json:"data"`
			Error interface{} `json:"error"`
		} `json:"display"`
		Content struct {
			DataType string `json:"dataType"`
			Type     string `json:"type"`
			Fields   struct {
				Description string `json:"description"`
				Id          struct {
					Id string `json:"id"`
				} `json:"id"`
				Metadata string `json:"metadata"`
				Name     string `json:"name"`
				Symbol   string `json:"symbol"`
				Url      string `json:"url"`
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
}

func ObjectsPrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}
	var res Object
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}

	fmt.Println(res.Data.Owner.AddressOwner)

}
