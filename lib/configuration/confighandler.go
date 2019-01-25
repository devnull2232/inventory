package configuration;

import (
	"fmt"
	"github.com/devnull2232/inventory/lib/persistence/dblayer"
	"encoding/json"
	"os"
)

type InventoryConfig struct{
	DB *DatabaseConfig `json:"db_config"`
}

type DatabaseConfig struct{
	DBType			dblayer.DBType	`json:"db_type"`
	DBConnection	string			`json:"db_connection"`
}

var (
	DBTypeDefault = dblayer.DBType("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
)

func ExtractConfiguration(filename string) (InventoryConfig, error){
	conf := InventoryConfig{
		DB: &DatabaseConfig{
			DBType: DBTypeDefault,
			DBConnection: DBConnectionDefault,
		},
	}
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Configuration file not found. Continuing with default values")
		return conf, err
	}
	json.NewDecoder(file).Decode(&conf)

	if v := os.Getenv("MONGO_URL"); v != ""{
		conf.DB.DBType =  "mongodb"
		conf.DB.DBConnection = v
	}
	return conf, nil
}