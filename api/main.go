package main

import (
	"flag"
	"fmt"
	"github.com/kwanok/spatial-query-study/api/config"
	"github.com/kwanok/spatial-query-study/api/db"
	"github.com/kwanok/spatial-query-study/api/server"
	"github.com/spf13/viper"
)

var (
	ConfigFilePath = flag.String("config", "runtime-config-local.yaml", "Path to config file")
)

func init() {
	flag.Parse()

	viper.AddConfigPath("./config")
	viper.SetConfigFile(*ConfigFilePath)
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
