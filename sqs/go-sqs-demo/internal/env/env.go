package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const sqsQueueURLEnv = "SQS_QUEUE_URL"

func LoadSQSQueueURL() string {
	_ = godotenv.Load()

	queueURL := os.Getenv(sqsQueueURLEnv)
	if queueURL == "" {
		log.Fatalf("%s environment variable is required", sqsQueueURLEnv)
	}

	return queueURL
}
