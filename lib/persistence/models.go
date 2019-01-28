package persistence;

import (
	"time"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Item struct {
	ID			*primitive.ObjectID		`json:"id" bson:"_id"`
	Tags		[]string				`json:"tags" bson:"tags"`
	ProductName	string					`json:"productName" bson:"productName"`
	Image		string					`json:"image" bson:"image"`
	Size		string					`json:"size" bson:"size"`
	StockState	StockState				`json:"stockState" bson:"stockState"`
}

type StockState struct {
	OriginalStockLevel			int32				`json:"originalLevel" bson:"originalLevel"`
	CurrentStockLevel			int32				`json:"currentLevel" bson:"currentLevel"`
	LastRestock					time.Time			`json:"lastRestock" bson:"lastRestock"`
	LastPrice					float64				`json:"lastPrice" bson:"lastPrice"`
	StockWaitingToBeReceived	int32				`json:"stockToReceive" bson:"stockToReceive"`
	JustificationStruct			JustificationStruct	`json:"justificationStruct" bson:"justificationStruct"`
}



type Justification struct {
	Text		string		`json:"text" bson:"text"`
	Employee	Employee	`json:"employee" bson:"employee"`
	PostedOn	time.Time	`json:"postedOn" bson:"postedOn"`
}

type JustificationStruct struct {
	LastJustification	Justification	`json:"lastJustification" bson:"lastJustification"`
	EditHistory			[]Justification	`json:"editHistory" bson:"editHistory"`
}

type Employee struct {
	Name			string		`json:"name" bson:"name"`
	Email			string		`json:"email" bson:"email"`
	EmployeeCode	string		`json:"employeeCode" bson:"employeeCode"`
	PhoneNumber		[]PhoneNumber	`json:"phoneNumber" bson:"phoneNumber"`
}

type PhoneNumber struct {
	CountryCode		int32		`json:"countryCode" bson:"countryCode"`
	Number			string		`json:"number" bson:"number"`
	PhoneType		PhoneType	`json:"phoneType" bson:"phoneType"`
}

type PhoneType int32

const (
	WORK PhoneType = iota + 1
	HOME
	CELLPHONE
)