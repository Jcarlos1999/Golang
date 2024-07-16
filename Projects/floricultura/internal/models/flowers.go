package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Flower struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Price       float64
	HarvestDate string
}

func NewFlower(name string, price float64, harvest string) (*Flower, error) {

	flower := &Flower{
		Name:        name,
		Price:       price,
		HarvestDate: harvest,
	}

	err := flower.Validade()
	if err != nil {
		return nil, err
	}

	return flower, nil
}

func (f *Flower) Validade() error {

	if f.Name == "" {
		return errors.New("Error: Flower should be have a name")
	}
	if f.Price <= 0 {
		return errors.New("Error: price must be greater than 0")
	}
	if f.HarvestDate == "" {
		return errors.New("Error: harvest time is not valid")
	}

	return nil
}

func (f *Flower) ToString() {
	fmt.Printf("Flower:%s,\nPrice:%0.2f,\nHarvest:%s", f.Name, f.Price, f.HarvestDate)
}
