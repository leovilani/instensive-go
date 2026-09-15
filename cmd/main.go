package main

// go run main.go - Command to run the program
// go build main.go - Command to build the program into an executable
// go mod init github.com/yourusername/yourproject - Command to initialize a new Go module

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/leovilani/go-study/internal/order/infra/database"
	"github.com/leovilani/go-study/internal/order/usecase"
	"github.com/leovilani/go-study/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	maxWorkers := 3
	// with a wait group I can manage my go routine workers
	wg := sync.WaitGroup{}
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUsecase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	// add in a wait group 3 threads
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, uc, i)
	}
	// will wait all workes run before finish the app
	wg.Wait()

	// this was before RabbitMQ, now the data comes from RabbitMQ
	// input := usecase.OrderInputDTO{
	// 	ID:    "1234",
	// 	Price: 100,
	// 	Tax:   10,
	// }
	// output, err := uc.Execute(input)
	// if err != nil {
	// 	panic(err)
	// }
	// println(output.FinalPrice)
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUsecase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}
		input.Tax = 10.0
		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}
		msg.Ack(false)
		println("Worker", workerId, "processed order", input.ID)
	}
}
