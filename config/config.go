package config

import "os"

type Configuration struct {
	Port           string
	DriverName     string
	DataSourceName string
}

var Config Configuration

func Initialize() {
	Config = Configuration{
		Port : os.Getenv("Port"),
		DriverName:  = os.Getenv("DriverName"),
		DataSourceName = os.Getenv("DataSourceName"),
	}
}
