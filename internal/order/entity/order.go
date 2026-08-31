package entity

import (
	"errors"
)

// In Go, interfaces are implemented implicitly.
// A type satisfies an interface by implementing all of its methods.
type OrderRepositoryInterface interface {
	Save(order *Order) error
}

// struct - Are used to define custom data types in Go. They are a collection of fields that can hold different types of data. Structs are used to create complex data structures and to model real-world entities in programming.
type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	// & is the pointer (ponteiro)
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	// Go dont have try catch, u have to verify and return error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// IsValid now is part of the Order struct
// IsValid with I is a public methos, if it starts with a lowercase letter, it is a private method and can only be accessed within the same package.
func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("Invalid ID")
	}
	if o.Price == 0 {
		return errors.New("Invalid price")
	}
	if o.Tax == 0 {
		return errors.New("Invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
