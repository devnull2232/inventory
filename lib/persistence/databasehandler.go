package persistence;

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type DatabaseHandler interface{
	GetAllItems(*context.Context) (*[]Item, error)
	GetClient() (*mongo.Client)
	GetItemsByTag(*context.Context, ...string) (*[]Item, error)
	GetItemById(*context.Context, *primitive.ObjectID) (*Item, error)
	// CreateItem(Item) (*Item, error)
	// GetStockHistory(*[]StockState, error)
}