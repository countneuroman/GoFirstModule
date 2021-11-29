package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Input name(s)")
	fmt.Println("Input q to display greetings ")
	var names []string
	for {
		reader := bufio.NewReader(os.Stdin)
		input, inputError := reader.ReadString('\n')

		if inputError != nil {
			log.Fatal(inputError)
		}

		if input == "q\r\n" {
			break
		}

		names = append(names, input)
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}

	fmt.Scanf("h")
}
