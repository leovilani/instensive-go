package main

// go run main.go - Command to run the program
// go build main.go - Command to build the program into an executable
// go mod init github.com/yourusername/yourproject - Command to initialize a new Go module

import (
	"database/sql"

	"github.com/leovilani/go-study/internal/order/infra/database"
	"github.com/leovilani/go-study/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUsecase(repository)

	defer db.Close()

	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100,
		Tax:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output.FinalPrice)
}
