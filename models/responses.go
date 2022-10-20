package models

type Response_restaurants struct {
	Restaurant User
	Address    Address
	Distance   float64
}

type Response_cart struct {
	ProductName string
	Entity      int
}
