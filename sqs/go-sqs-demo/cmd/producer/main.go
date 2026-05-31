package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/bedoodev/go-sqs-demo/internal/env"
	"github.com/bedoodev/go-sqs-demo/internal/model"
	"github.com/bedoodev/go-sqs-demo/internal/producer"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("error while reading config %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	orderProducer := producer.New(sqsClient, env.LoadSQSQueueURL())

	randomEventID, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("error creating random uuid %v", err)
	}
	randomOrderID := rand.IntN(999999)
	randomUserID := rand.IntN(999999)
	randomAmount := rand.Float64() * 10000
	currencies := []string{"TRY", "USD", "EUR", "GBP"}
	randomCurrency := rand.IntN(len(currencies) - 1)

	event := model.OrderCreatedEvent{
		EventID:    randomEventID.String(),
		EventType:  "order.created",
		OccurredAt: time.Now(),
		Data: model.OrderData{
			OrderID:  fmt.Sprintf("ORD-%d", randomOrderID),
			UserID:   fmt.Sprintf("USR-%d", randomUserID),
			Amount:   randomAmount,
			Currency: currencies[randomCurrency],
		},
	}

	err = orderProducer.SendOrder(ctx, &event)
	if err != nil {
		log.Fatal(err)

	}
}
