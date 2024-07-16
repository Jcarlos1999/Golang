package models

import (
	"errors"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Items      []Flower `gorm:"many2many:Order_Flower;"`
	FinalPrice float64
}

func NewOrder(item []Flower, finalPrice float64) (*Order, error) {

	order := &Order{
		Items:      item,
		FinalPrice: finalPrice,
	}
	err := order.Validade()
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (o *Order) Validade() error {

	for _, i := range o.Items {

		if i.Name == "" {
			return errors.New("Error: Invalid flower")
		}
		if i.Price < 0 {
			return errors.New("Error: Final should be greater than 0")
		}
	}

	return nil
}
