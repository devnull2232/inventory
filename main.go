package main

import (
	"flag"
	"fmt"
	"github.com/devnull2232/inventory/lib/configuration"
	"github.com/devnull2232/inventory/lib/persistence/dblayer"
	//"github.com/devnull2232/inventory/lib/persistence/mongolayer"
	"context"
	// "time"
	// "log"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/davecgh/go-spew/spew"
)

func main(){
	// Initial configuration:
	// receive configuration path from terminal
	confPath := flag.String("conf", `lib/configuration/config.json`, "this sets the path to the configuration file")
	flag.Parse()

	// store configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println(config.DB)
	fmt.Println("Connecting to database cluster...")
	//test

	ctx := context.Background()
	mongohandler, err := dblayer.NewPersistenceLayer(config.DB.DBType, config.DB.DBConnection, &ctx)
	if err != nil{
		panic(err)
	}
	
	// items, err := mongohandler.GetAllItems(&ctx)
	// items, err := mongohandler.GetItemsByTag(&ctx, "gaseosa", "xd")
	id, err := primitive.ObjectIDFromHex("5c4f5252a2ac9e3ed75c2424")
	item, err := mongohandler.GetItemById(&ctx, &id)
	if err != nil{
		panic(err)
	}

	//fmt.Printf("%v\n", items)
	 spew.Dump(*item)
}