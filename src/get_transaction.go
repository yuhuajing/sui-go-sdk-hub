package src

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func GetTransction(hash []string) {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)
	if len(hash) == 0 {
		return
	} else if len(hash) == 1 {
		rsp, err := cli.SuiGetTransactionBlock(ctx, models.SuiGetTransactionBlockRequest{
			Digest: hash[0],
			// only fetch the effects field
			Options: models.SuiTransactionBlockOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowBalanceChanges: true,
				ShowObjectChanges:  true,
			},
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		utils.PrettyPrint(rsp)
		return
	}
	digests := make([]string, len(hash))
	for _, v := range hash {
		digests = append(digests, v)
	}
	// fetch multiple transactions in one batch request
	rsp2, err := cli.SuiMultiGetTransactionBlocks(ctx, models.SuiMultiGetTransactionBlocksRequest{
		Digests: digests,
		// only fetch the effects field
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, transactionBlock := range rsp2 {
		utils.PrettyPrint(*transactionBlock)
	}

}

type T struct {
	Digest      string `json:"digest"`
	Transaction struct {
		Data struct {
			MessageVersion string `json:"messageVersion"`
			Transaction    struct {
				Kind   string `json:"kind"`
				Inputs []struct {
					Type      string `json:"type"`
					Value     string `json:"value"`
					ValueType string `json:"valueType"`
				} `json:"inputs"`
				Transactions []struct {
					MoveCall struct {
						Arguments []struct {
							Input int `json:"Input"`
						} `json:"arguments"`
						Function string `json:"function"`
						Module   string `json:"module"`
						Package  string `json:"package"`
					} `json:"MoveCall"`
				} `json:"transactions"`
			} `json:"transaction"`
			Sender  string `json:"sender"`
			GasData struct {
				Payment []struct {
					ObjectId string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"payment"`
				Owner  string `json:"owner"`
				Price  string `json:"price"`
				Budget string `json:"budget"`
			} `json:"gasData"`
		} `json:"data"`
		TxSignatures []string `json:"txSignatures"`
	} `json:"transaction"`
	RawTransaction string `json:"rawTransaction"`
	Effects        struct {
		MessageVersion string `json:"messageVersion"`
		Status         struct {
			Status string `json:"status"`
		} `json:"status"`
		ExecutedEpoch string `json:"executedEpoch"`
		GasUsed       struct {
			ComputationCost         string `json:"computationCost"`
			StorageCost             string `json:"storageCost"`
			StorageRebate           string `json:"storageRebate"`
			NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
		} `json:"gasUsed"`
		ModifiedAtVersions []struct {
			ObjectId       string `json:"objectId"`
			SequenceNumber string `json:"sequenceNumber"`
		} `json:"modifiedAtVersions"`
		SharedObjects     interface{} `json:"sharedObjects"`
		TransactionDigest string      `json:"transactionDigest"`
		Created           interface{} `json:"created"`
		Mutated           []struct {
			Owner struct {
				AddressOwner string `json:"AddressOwner"`
			} `json:"owner"`
			Reference struct {
				ObjectId string `json:"objectId"`
				Version  int    `json:"version"`
				Digest   string `json:"digest"`
			} `json:"reference"`
		} `json:"mutated"`
		Deleted   interface{} `json:"deleted"`
		GasObject struct {
			Owner struct {
				AddressOwner string `json:"AddressOwner"`
			} `json:"owner"`
			Reference struct {
				ObjectId string `json:"objectId"`
				Version  int    `json:"version"`
				Digest   string `json:"digest"`
			} `json:"reference"`
		} `json:"gasObject"`
		EventsDigest string   `json:"eventsDigest"`
		Dependencies []string `json:"dependencies"`
	} `json:"effects"`
	Events []struct {
		Id struct {
			TxDigest string `json:"txDigest"`
			EventSeq string `json:"eventSeq"`
		} `json:"id"`
		PackageId         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		ParsedJson        struct {
			IsVerified bool `json:"is_verified"`
		} `json:"parsedJson"`
		Bcs         string `json:"bcs"`
		TimestampMs string `json:"timestampMs"`
	} `json:"events"`
	TimestampMs string `json:"timestampMs"`
	Checkpoint  string `json:"checkpoint"`
}
