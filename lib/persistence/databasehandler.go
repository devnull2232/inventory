package persistence;

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"

)

type DatabaseHandler interface{
	GetAllItems(*context.Context) ([]Item, error)
	GetClient() (*mongo.Client)
	// GetItemsByTag(tag string) (*[]Item, error)
	// GetItemById(id string) (*Item, error)
	// CreateItem(Item) (*Item, error)
	// GetStockHistory(*[]StockState, error)
}