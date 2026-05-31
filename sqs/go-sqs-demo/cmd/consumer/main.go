package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/bedoodev/go-sqs-demo/internal/consumer"
	"github.com/bedoodev/go-sqs-demo/internal/env"
)

func main() {

	// Load AWS Config from ~/.aws/credentials
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("An error occurred while trying to connect to AWS. %v", err)
	}

	// Create SQS Client
	sqsClient := sqs.NewFromConfig(cfg)

	orderConsumer := consumer.New(sqsClient, env.LoadSQSQueueURL())

	for {
		messages, err := orderConsumer.ReceiveMessages(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, message := range messages {
			event, err := consumer.ParseMessage(message)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("%+v\n", event)

			err = orderConsumer.DeleteMessage(ctx, *message.ReceiptHandle)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

}
