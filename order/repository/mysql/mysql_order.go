package mysql

import (
	"assignment-dua-fga/models"
	"fmt"
	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	DB *gorm.DB
}

func (m mysqlOrderRepository) Fetch() (res []models.Order, err error) {
	orders := make([]models.Order, 0)
	err = m.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (m mysqlOrderRepository) Store(order models.Order) error {
	err := m.DB.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (m mysqlOrderRepository) Update(id int, order models.Order) error {
	fmt.Println(order)
	err := m.DB.Transaction(
		func(tx *gorm.DB) error {
			err := m.DB.Model(&order).Where("id = ?", id).Updates(&order).Error
			if err != nil {
				return err
			}
			err = m.DB.Model(&order).Association("Items").Replace(order.Items)
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func (m mysqlOrderRepository) Delete(id int) error {
	err := m.DB.Delete((&models.Order{}), id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMysqlOrderRepository(Db *gorm.DB) models.OrderRepository {
	return &mysqlOrderRepository{DB: Db}
}
