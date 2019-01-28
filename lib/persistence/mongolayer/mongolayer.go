package mongolayer

import (
	"github.com/devnull2232/inventory/lib/persistence"
	// mgo "gopkg.in/mgo.v2"
	// migrating from mgo.v2 to the mongodb official driver
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/mongodb/mongo-go-driver/x/bsonx"
	"log"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"context"
	"time"
	"fmt"
)

const (
	DB = "inventory"	// db name
	CATALOG = "catalog" // collection name
	TIMEOUT = 5*time.Second
)

type MongoDBLayer struct{
	//session *mgo.Session
	client *mongo.Client
}

// mgo.v2
// func NewMongoDBLayer(connection string) (*MongoDBLayer, error){
// 	s, err := mgo.Dial(connection)
// 	if err != nil{
// 		return nil, err
// 	}
// 	return &MongoDBLayer{
// 		session: s,
// 	}, err
// }

func (m *MongoDBLayer) GetClient() (*mongo.Client){
	return m.client
}

// This method creates a new client to connect to the cluster specified by 'connection'
func NewMongoDBLayer(connection string) (*MongoDBLayer, error){
	// Creating client
	client, err := mongo.NewClient(connection)
	if err != nil{
		return nil, err
	}
	return &MongoDBLayer{
		client: client,
	}, err
}

// mgo.v2
// func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session{
// 	return mgoLayer.session.Copy()
// }

// This method initializes the client by starting background monitoring goroutines
// It returns a pointer to a context.Context that has 'timeout', a context.CancelFunc, and a error
func (mongoLayer *MongoDBLayer) InitializeClient(ctx *context.Context) error {
	err := mongoLayer.client.Connect(*ctx)
	if err != nil{
		log.Fatal("Could not connect to cluster. Error: ", err)
	}
	err = mongoLayer.client.Ping(*ctx, nil)
	if err == nil {
		fmt.Println("Connected to the Database!")

	}
	return err
}

// TODO: implement DatabaseHandler interface


// mgo.v2
// // Gets all items in the catalog
// func (mgoLayer *MongoDBLayer) GetAllItems() (*[]persistence.Item, error){
// 	s := mgoLayer.getFreshSession()
// 	defer s.Close()
// 	items := []persistence.Item{}
// 	err := s.DB(DB).C(CATALOG).Find(nil).All(&items)
// 	return &items, err
// }

// Gets all items in the catalog
// mongoLayer MUST call this method ONLY AFTER having called InitializeClient
func (mongoLayer *MongoDBLayer) GetAllItems(ctx *context.Context) (*[]persistence.Item, error){
	return mongoLayer.GetItemsByTag(ctx)
}

func traverseItems(cursor *mongo.Cursor, ctx *context.Context, items *[]persistence.Item) error{
	var (
		err error
		item persistence.Item
	)
	for (*cursor).Next(*ctx) {
		err = (*cursor).Decode(&item)
		if err != nil{
			return err
		}
		(*items) = append(*items, item)
	}
	return err
}

func (mongoLayer *MongoDBLayer) GetItemsByTag(ctx *context.Context, strings ...string) (*[]persistence.Item, error){
	items, coll, tags := []persistence.Item{}, mongoLayer.client.Database(DB).Collection(CATALOG), bson.A{}
	for _, s := range strings {
		tags = append(tags, s)
	}
	var (
		cursor mongo.Cursor
		err error
	)
	if len(tags) != 0 {
		cursor, err = coll.Find(*ctx, bson.D{
			{"tags", bson.D{{"$all", tags}}},
		})
	} else {
		cursor, err = coll.Find(*ctx, bson.D{})
	}
	defer cursor.Close(*ctx)
	err = traverseItems(&cursor, ctx, &items)
	if err != nil{
		return nil, err
	}
	return &items, err
}

func (mongoLayer *MongoDBLayer) GetItemById(ctx *context.Context, id *primitive.ObjectID) (*persistence.Item, error){
	item, coll := persistence.Item{}, mongoLayer.client.Database(DB).Collection(CATALOG)
	result := coll.FindOne(*ctx, bson.D{{"_id", id}})
	err := result.Decode(&item)
	if err != nil{
		return nil, err
	}
	return &item, err 
}

// mgo.v2
// func (mgoLayer *MongoDBLayer) GetItemsByTag(tag string) (*[]persistence.Item, error){}
// func (mgoLayer *MongoDBLayer) GetItemById(id string) (*persistence.Item, error){}
// func (mgoLayer *MongoDBLayer) CreateItem(persistence.Item) (*persistence.Item, error){}
// func (mgoLayer *MongoDBLayer) GetStockHistory(*[]persistence.StockState, error){}

// TODO: 
// TODO: Add sections.
// TODO: Add items to a section.
// TODO: Update an item. Add more tags, add sizes to choose from, change quatity, etc.
// TODO: Get a number of items
// TODO: Get a section of items (Sections and tags are different)
// A section would be something like drinks, fruits, spices, etc.
// Tags would be something like $, $$ or contains alcohol, perishable, etc.