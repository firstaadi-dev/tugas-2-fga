package models

type Item struct {
	ID          int
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}
