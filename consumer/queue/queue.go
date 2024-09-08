package queue

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	serviceSqs *sqs.SQS
	queueUrl   *sqs.GetQueueUrlOutput
)

func StartQueue() {
	fmt.Println("Criando aws session")

	awsSession := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String("http://localhost:4566"),
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("test", "test", "test"),
	}))
	fmt.Println("Seção criada")

	serviceSqs = sqs.New(awsSession)

	queueName := "queue1"
	queueUrl, _ = serviceSqs.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
}

func ReceveidMessages() *sqs.ReceiveMessageOutput {
	msgResult, err := serviceSqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl.QueueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(30),
	})

	if err != nil {
		panic(err)
	}

	return msgResult
}

func DeleteMessage(receiptHandler *string) {
	_, err := serviceSqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl.QueueUrl,
		ReceiptHandle: receiptHandler,
	})

	if err != nil {
		panic(err)
	}
}
