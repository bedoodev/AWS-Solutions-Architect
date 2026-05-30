# Go SQS Demo

A simple Producer-Consumer application written in Go using Amazon SQS.

## Overview

This project demonstrates the basic usage of Amazon SQS with Go.

The application consists of two separate components:

- Producer: Creates an order event and sends it to an SQS queue.
- Consumer: Reads messages from the queue, deserializes the event, processes it, and deletes the message.

The purpose of this project is to understand the fundamentals of SQS message flow rather than building a production-ready messaging system.

## Architecture

```plaintext
+------------+         +-------------+         +------------+
|  Producer  | ------> | Amazon SQS  | ------> |  Consumer  |
+------------+         +-------------+         +------------+
```

## Project Structure

```plaintext
go-sqs-demo/
├── cmd/
│   ├── producer/
│   │   └── main.go
│   └── consumer/
│       └── main.go
├── internal/
│   ├── consumer/
│   │   └── consumer.go
│   ├── producer/
│   │   └── producer.go
│   └── model/
│       └── event.go
├── go.mod
└── go.sum
```

## Event Model

The producer publishes an `OrderCreatedEvent`.

Example payload:

```json
{
  "event_id": "019e79cd-1c8d-74c2-bdec-45c4d557331e",
  "event_type": "order.created",
  "occurred_at": "2026-05-30T16:52:25Z",
  "data": {
    "order_id": "ORD-1",
    "user_id": "USR-1",
    "amount": 349.99,
    "currency": "TRY"
  }
}
```

## Producer Flow

1. Load AWS configuration.
2. Create an SQS client.
3. Generate a UUID v7 event ID.
4. Create an `OrderCreatedEvent`.
5. Serialize the event into JSON.
6. Send the message to Amazon SQS.

## Consumer Flow

1. Load AWS configuration.
2. Create an SQS client.
3. Poll SQS for messages.
4. Deserialize the JSON payload into an `OrderCreatedEvent`.
5. Execute business logic.
6. Delete the message from the queue.

## Running the Producer

```bash
go run ./cmd/producer
```

## Running the Consumer

```bash
go run ./cmd/consumer
```

## AWS Configuration

```bash
aws configure
```

The AWS SDK v2 automatically loads credentials and region information from:

```plaintext
~/.aws/credentials
~/.aws/config
```

## Concepts Demonstrated

- Amazon SQS message publishing
- Amazon SQS message consumption
- Long polling
- Message serialization and deserialization
- Message deletion
- UUID v7 event generation
- Basic producer-consumer architecture

## Limitations

This repository is intentionally kept simple and focuses only on the core SQS workflow.

The following production-ready features have intentionally been omitted:

- Graceful shutdown
- Worker pools
- Concurrent message processing
- Retry strategies
- Dead Letter Queues (DLQ)
- Structured logging
- Metrics and monitoring
- Distributed tracing
- Health checks
- Configuration management
- Idempotency handling
- Visibility timeout extensions

The goal is to provide a minimal and easy-to-understand introduction to Amazon SQS using Go.
