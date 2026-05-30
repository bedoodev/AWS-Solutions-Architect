package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/bedoodev/go-sqs-demo/internal/model"
)

type OrderProducer struct {
	client   *sqs.Client
	queueURL string
}

func New(client *sqs.Client, queueURL string) *OrderProducer {
	return &OrderProducer{client: client, queueURL: queueURL}
}

func (p *OrderProducer) SendOrder(ctx context.Context, event *model.OrderCreatedEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal order created event: %w", err)
	}

	_, err = p.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(p.queueURL),
		MessageBody: aws.String(string(payload)),
	})
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
