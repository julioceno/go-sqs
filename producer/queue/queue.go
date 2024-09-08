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

func SendMessage(message string) {
	_, err := serviceSqs.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(message),
		QueueUrl:     queueUrl.QueueUrl,
		DelaySeconds: aws.Int64(10),
	})

	if err != nil {
		panic(err)
	}
}
