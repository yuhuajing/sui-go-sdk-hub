package src

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
)

func GetAllCoinValues() {
	var ctx = context.Background()

	rsp, err := cli.SuiXGetAllBalance(ctx, models.SuiXGetAllBalanceRequest{
		Owner: "0x2cba318afda7b04022cf7522925e11117a966b35771bc9e10679c5b0a0c7739b",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	AllCoinsPrettyPrint(rsp)
}

type CoinStruct struct {
	CoinType        string `json:"coinType"`
	CoinObjectCount int    `json:"coinObjectCount"`
	TotalBalance    string `json:"totalBalance"`
	LockedBalance   struct {
		EpochId int `json:"epochId"`
		Number  int `json:"number"`
	} `json:"lockedBalance"`
}

func AllCoinsPrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}
	var res []CoinStruct
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	for _, v := range res {
		fmt.Println(v.CoinType)
	}
}
