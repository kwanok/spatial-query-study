package main

import (
	"fmt"
	"github.com/kwanok/spatial-query-study/api/config"
	"github.com/kwanok/spatial-query-study/api/db"
	"github.com/kwanok/spatial-query-study/api/server"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigFile("../global_config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config.RuntimeConf)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.RuntimeConf)
}

func main() {
	db.Start()
	server.Start()

	defer db.Conn.Close()
}
