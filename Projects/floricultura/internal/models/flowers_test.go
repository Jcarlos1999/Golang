package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsOkIfIDIsCorrect(t *testing.T) {

	flower := Flower{Name: "Orchid", Price: 23}
	assert.Error(t, flower.Validade(), "Error: invalid Name")
}

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {

	flower := Flower{Price: 23}
	assert.Error(t, flower.Validade(), "Error: invalid price")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {

	flower := Flower{Name: "Orchid"}
	assert.Error(t, flower.Validade(), "Error: pric must be greater than 0")
}

func TestIfItGetsAnErrorIfHarvestIsBlank(t *testing.T) {

	flower := Flower{Name: "Orchid", Price: 23}
	assert.Error(t, flower.Validade(), "")
}

func TestIfItDoingFine(t *testing.T) {
	flower := Flower{Name: "Orchid", Price: 23, HarvestDate: "2023-01-02"}
	assert.Equal(t, flower.Validade(), nil)
}

func TestIfItNewFlowerNotReturnError(t *testing.T) {
	_, err := NewFlower("Orchid", 10, time.Now().String())

	assert.Equal(t, err, nil)
}

func TestIfItNewFlowerReturnErrorAtID(t *testing.T) {
	_, err := NewFlower("Orchid", 10, time.Now().String())

	assert.Error(t, err, "Error: invalid type ID")
}

func TestIfItNewFlowerReturnErrorAtPrice(t *testing.T) {
	_, err := NewFlower("Orchid", -10, time.Now().String())

	assert.Error(t, err, "Error: price must be greater than 0")
}

func TestIfItNewFlowerReturnErrorAtHarvestDate(t *testing.T) {
	_, err := NewFlower("Orchid", -10, time.Now().String())

	assert.Error(t, err, "Erro: price must be greater than 0")
}
