package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kwanok/spatial-query-study/api/config"
	"log"
)

var Conn *sql.DB

func NewDatasource() *config.Datasource {
	datasource := &config.RuntimeConf.Datasource
	return &config.Datasource{
		Host:     datasource.Host,
		Port:     datasource.Port,
		Database: datasource.Database,
		User:     datasource.User,
		Password: datasource.Password,
	}
}

func Start() {
	db, err := sql.Open("mysql", NewDatasource().GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	Conn = db

	err = Conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
