package config

import "fmt"

var RuntimeConf = Api{}

type Datasource struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (d *Datasource) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}

type Server struct {
	Port string `yaml:"port"`
}

type Api struct {
	Server     Server     `yaml:"server"`
	Datasource Datasource `yaml:"datasource"`
}
