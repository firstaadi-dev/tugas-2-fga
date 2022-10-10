package usecase

import "assignment-dua-fga/models"

type OrderUsecase struct {
	OrderRepository models.OrderRepository
}

func (o OrderUsecase) Fetch() ([]models.Order, error) {
	res, err := o.OrderRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o OrderUsecase) Store(order models.Order) error {
	err := o.OrderRepository.Store(order)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderUsecase) Update(id int, order models.Order) error {
	err := o.OrderRepository.Update(id, order)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderUsecase) Delete(id int) error {
	err := o.OrderRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewOrderUsecase(a models.OrderRepository) models.OrderUsecase {
	return &OrderUsecase{OrderRepository: a}
}
