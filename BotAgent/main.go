package main

import (
	"botagent/bots"
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Language: ")

	var language string
	_, err := fmt.Scanln(&language)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bot, err := createBot(language)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	answer := ""
	for {
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		a, isExit := bot.ParseUserFrase(answer)
		fmt.Println(a)
		if isExit {
			return
		}
	}
}

func createBot(language string) (bots.Bot, error) {

	var b bots.Bot

	switch language {
	case "English":
		b = bots.BotEnglish{Name: "John"}
	case "Russian":
		b = bots.BotRussian{Name: "Иван"}
	default:
		return b, errors.New(bots.LanguageErrorText)
	}
	return b, nil
}
