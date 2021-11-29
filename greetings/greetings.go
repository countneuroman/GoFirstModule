package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func Hello(name string) (string, error) {
	if name == "\r\n" || name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(RandomFormat(), name)

	return message, nil
}

func RandomFormat() string {
	formats := []string{
		"Hi, %v! Welcome!",
		"Great to see you %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
