# Introduction to SQS

## What is Amazon SQS?

Amazon Simple Queue Service (SQS) is a fully managed message queue service that enables decoupling and asynchronous communication between distributed applications.

### Why Use SQS?

Without SQS:

```plaintext
Application A ---> Application B
```

With SQS:

```plaintext
Application A ---> SQS Queue ---> Application B
```

Benefits:

- Loose coupling between services
- Independent scaling
- Improved reliability
- Asynchronous processing
- Reduced system dependencies

---

## Core Components

### Producer

A producer sends messages to an SQS queue.

```plaintext
Producer ---> Queue
```

### Queue

A queue temporarily stores messages until they are processed.

### Consumer

A consumer receives and processes messages.

```plaintext
Queue ---> Consumer
```

---

## Queue Types

### Standard Queue

Characteristics:

- Nearly unlimited throughput
- At-least-once delivery
- Best-effort ordering

Use Cases:

- Background jobs
- Event-driven systems
- Analytics pipelines
- Log processing

---

### FIFO Queue

FIFO = First In First Out

Characteristics:

- Exactly-once processing
- Strict ordering
- Lower throughput than Standard Queue

Requirements:

- MessageGroupId
- Optional MessageDeduplicationId

Use Cases:

- Payments
- Banking systems
- Inventory management

---

## Message Lifecycle

```plaintext
Producer
    |
    v
Send Message
    |
    v
SQS Queue
    |
    v
Receive Message
    |
    v
Process Message
    |
    v
Delete Message
```

Messages are not removed automatically after being received.

DeleteMessage must be called after successful processing.

---

## Visibility Timeout

When a consumer receives a message:

- Message becomes invisible
- Other consumers cannot see it

```plaintext
Receive Message
        |
        v
Visibility Timeout
        |
        +--> DeleteMessage() => Removed
        |
        +--> Consumer Failure => Visible Again
```

---

## At-Least-Once Delivery

Standard Queues provide:

- At-Least-Once Delivery

Meaning:

A message can be delivered multiple times.

Applications should therefore be idempotent.

Examples:

- UPSERT instead of INSERT
- Event deduplication
- Unique event IDs

---

## Long Polling

Without Long Polling:

```go
ReceiveMessage()
```

With Long Polling:

```go
ReceiveMessageInput{
    WaitTimeSeconds: 20,
}
```

Benefits:

- Fewer API calls
- Lower cost
- Reduced empty responses

---

## Message Retention

Messages remain in a queue even if no consumer processes them.

Retention Period:

- Minimum: 1 minute
- Maximum: 14 days
- Default: 4 days

---

## Dead Letter Queue (DLQ)

A DLQ stores messages that repeatedly fail processing.

```plaintext
Producer
    |
    v
Main Queue
    |
    +--> Success
    |
    +--> Failure
            |
            v
           DLQ
```

Benefits:

- Isolates problematic messages
- Simplifies debugging
- Prevents processing bottlenecks

---

## Delay Queues

Delay Queues postpone message visibility.

Example:

```plaintext
Send Message
      |
      v
Wait 30 Seconds
      |
      v
Consumer Receives Message
```

Use Cases:

- Scheduled processing
- Retry workflows
- Delayed notifications

---

## Message Size Limit

Maximum message size:

```plaintext
256 KB
```

For larger payloads:

- Store data in S3
- Send S3 object references through SQS

---

## Security

### IAM Policies

Control access to:

- SendMessage
- ReceiveMessage
- DeleteMessage
- CreateQueue
- PurgeQueue

---

### Encryption

#### SSE-SQS

AWS-managed encryption.

- Enabled by default
- No additional setup

#### SSE-KMS

KMS-managed encryption.

Benefits:

- Compliance support
- Auditability
- Fine-grained access control

Drawback:

- Additional KMS request charges

---

## Monitoring

CloudWatch Metrics:

- ApproximateNumberOfMessagesVisible
- ApproximateNumberOfMessagesNotVisible
- NumberOfMessagesSent
- NumberOfMessagesReceived
- NumberOfMessagesDeleted
- ApproximateAgeOfOldestMessage

---

## Common Architectures

### Worker Queue Pattern

```plaintext
Producer
    |
    v
SQS
    |
    +--> Worker 1
    +--> Worker 2
    +--> Worker 3
```

### Event-Driven Architecture

```plaintext
Order Service
      |
      v
Order Created Event
      |
      v
SQS
      |
      +--> Email Service
      +--> Analytics Service
      +--> Notification Service
```

---

## Important AWS Exam Facts

- Standard Queue = At-Least-Once Delivery
- FIFO Queue = Exactly-Once Processing
- Maximum Message Size = 256 KB
- Retention Period = 1 Minute - 14 Days
- Messages must be explicitly deleted
- Visibility Timeout prevents concurrent processing
- Long Polling reduces cost
- DLQ stores failed messages
- SQS is region-scoped
- SQS is fully managed

---

## Summary

Amazon SQS is a fully managed message queue service used to decouple applications and enable asynchronous communication.

Key Benefits:

- Scalability
- Reliability
- Durability
- Loose Coupling
- Simplicity

SQS is one of the foundational services used in event-driven and distributed architectures on AWS.
