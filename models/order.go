package models

import "time"

type Order struct {
	ID           int
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

type OrderUsecase interface {
	Fetch() ([]Order, error)
	Store(order Order) error
	Update(id int, order Order) error
	Delete(id int) error
}

type OrderRepository interface {
	Fetch() (res []Order, err error)
	Store(order Order) error
	Update(id int, order Order) error
	Delete(id int) error
}
