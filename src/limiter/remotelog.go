package limiter

import (
	"cloud.google.com/go/logging"
	"context"
	"log"
)

func initRemoteLog() *logging.Logger {
	projectID := "bitmovin-api"
	client, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}
	logName := "api-ratelimit-log"
	logger := client.Logger(logName)
	return logger
}
