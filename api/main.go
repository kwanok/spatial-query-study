package main

import (
	"github.com/kwanok/spatial-query-study/api/db"
	"github.com/kwanok/spatial-query-study/api/server"
	"github.com/spf13/viper"
)

func main() {
	defer db.Conn.Close()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	server.Start()
}
