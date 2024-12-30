package src

import (
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/joho/godotenv"
	"os"
)

var (
	cli sui.ISuiAPI
	key string
)

func init() {
	cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	_ = godotenv.Overload(".env")

	key = Env("PRIVATE_KEY", "")
}

func Env(key string, d string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return d
}
