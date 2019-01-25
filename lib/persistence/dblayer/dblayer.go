package dblayer

import (
	"github.com/devnull2232/inventory/lib/persistence"
	"github.com/devnull2232/inventory/lib/persistence/mongolayer"
	"context"
)

type DBType string;

const (
	MONGODB DBType = "mongodb"
)

func NewPersistenceLayer(dbtype DBType, connection string, ctx *context.Context) (persistence.DatabaseHandler, error){
	switch dbtype{
		case MONGODB:
			mongoDBHandler, err := mongolayer.NewMongoDBLayer(connection)
			if err != nil{
				return nil, err
			}
			err = mongoDBHandler.InitializeClient(ctx)
			return mongoDBHandler, err
	}
	return nil, nil
}