package entity_test

import (
	"testing"

	"github.com/leovilani/go-study/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}

	// test method without using a external library
	// if order.ID != "" {
	// 	t.Errorf("Expected empty ID, but got %s", order.ID)
	// }

	assert.Error(t, order.IsValid(), "Invalid ID")
}

func TestGivenAnEmptyPrice_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), "Invalid price")
}

func TestGivenAnEmptyTax_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 100}
	assert.Error(t, order.IsValid(), "Invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveAnOrder(t *testing.T) {
	order, err := entity.NewOrder("123", 100, 2)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
