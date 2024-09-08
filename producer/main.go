package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/julioceno/go-sqs/producer/queue"
)

func main() {
	queue.StartQueue()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Escreva sua mensagem: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		queue.SendMessage(input)
		fmt.Println("Enviado com sucesso")
	}
}
