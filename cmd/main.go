package main

// go run main.go - Command to run the program
// go build main.go - Command to build the program into an executable
// go mod init github.com/yourusername/yourproject - Command to initialize a new Go module

import (
	"database/sql"
	"fmt"

	"github.com/leovilani/go-study/internal/order/entity"
	"github.com/leovilani/go-study/internal/order/infra/database"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		panic(err)
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The final price is: %f", order.FinalPrice)

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	repository := database.NewOrderRepository(db)
	err = repository.Save(order)
	if err != nil {
		panic(err)
	}
}
