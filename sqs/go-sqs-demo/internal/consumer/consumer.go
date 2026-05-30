package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/bedoodev/go-sqs-demo/internal/model"
)

type OrderConsumer struct {
	client   *sqs.Client
	queueURL string
}

func New(client *sqs.Client, queueURL string) *OrderConsumer {
	return &OrderConsumer{client: client, queueURL: queueURL}
}

func (c *OrderConsumer) ReceiveMessages(ctx context.Context) ([]types.Message, error) {

	output, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(c.queueURL),
		WaitTimeSeconds: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("receive message: %w", err)
	}

	return output.Messages, nil
}

func ParseMessage(message types.Message) (*model.OrderCreatedEvent, error) {
	var event model.OrderCreatedEvent
	err := json.Unmarshal([]byte(*message.Body), &event)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &event, nil
}

func (c *OrderConsumer) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	})
	if err != nil {
		return fmt.Errorf("delete message: %w", err)
	}
	return nil
}
