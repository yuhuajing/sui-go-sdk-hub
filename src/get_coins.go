package src

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

func GetCoinMetaqData() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	rsp, err := cli.SuiXGetCoinMetadata(ctx, models.SuiXGetCoinMetadataRequest{
		CoinType: "0x2::sui::SUI",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	CoinsMetaDataPrettyPrint(rsp)
}

type CoinMetaData struct {
	Id          string `json:"id"`
	Decimals    int    `json:"decimals"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
}

func CoinsMetaDataPrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}
	var res CoinMetaData
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}

	fmt.Println(res.Name, res.Symbol, res.Decimals, res.Id, res.Description, res.IconUrl)

}
