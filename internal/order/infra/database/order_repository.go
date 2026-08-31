package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // the _ means that I am importing this lib but I dont will use it now
	"github.com/leovilani/go-study/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	// stmt means statement
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	// _ underscore means that I dont want the result of the query (something like: 1 line inserted)
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
