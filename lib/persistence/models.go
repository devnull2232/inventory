package persistence;

import (
	"time"
)

type Item struct {
	ID			string		`json:"id"`
	Tags		[]string	`json:"tags"`
	ProductName	string		`json:"product_name"`
	Image		string		`json:"image"`
	Size		string		`json:"size"`
	StockState	StockState	`json:"stock_state"`
}

type StockState struct {
	OriginalStockLevel			int32			`json:"original_level"`
	CurrentStockLevel			int32			`json:"current_level"`
	LastRestock					time.Time		`json:"last_restock"`
	LastPrice					float32			`json:"last_price"`
	StockWaitingToBeReceived	int32			`json:"stock_to_receive"`
	Justification				Justification	`json:"justification"`
}

type Justification struct {
	Text		string		`json:"text"`
	Employee	Employee	`json:"employee"`
	PostedOn	time.Time	`json:"posted_on"`
	EditHistory	EditHistory	`json:"edit_history"`
}

type Employee struct {
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	EmployeeCode	string		`json:"employee_code"`
	PhoneNumber		[]PhoneNumber	`json:"phone_number"`
}

type EditHistory struct {
	Text		string		`json:"text"`
	LastEdit	time.Time	`json:"last_edit"`
}

type PhoneNumber struct {
	CountryCode		int32		`json:"country_code"`
	Number			string		`json:"number"`
	PhoneType		PhoneType	`json:"phone_type"`
}

type PhoneType int32

const (
	WORK PhoneType = iota + 1
	HOME
	CELLPHONE
)