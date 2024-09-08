package main

import (
	"fmt"

	"github.com/julioceno/go-sqs/consumer/queue"
)

func main() {
	queue.StartQueue()

	for {
		result := queue.ReceveidMessages()
		for _, msg := range result.Messages {

			fmt.Println(*msg.Body)

			queue.DeleteMessage(msg.ReceiptHandle)
		}
	}
}
