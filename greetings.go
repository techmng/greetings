package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name should not be empty")
	}
	msg := fmt.Sprintf(randomFormatedMsg(), name)
	return msg, nil
}

func randomFormatedMsg() string {
	randomFormat := []string{
		"Hello %v, Welcome!",
		"Welcome Dear %v",
		"Your nice %v",
	}

	//return randomFormat[rand.Intn(len(randomFormat))]
	return randomFormat[0]
}
