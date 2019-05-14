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

	answer := 0
	for {
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		switch answer {
		case 1:
			fmt.Println(bot.SayHello())
		case 2:
			fmt.Println(bot.SayTime())
		case 3:
			fmt.Println(bot.SayDate())
		case 4:
			fmt.Println(bot.SayWeekday())
		case 5:
			fmt.Println(bot.SayBye())
			return
		default:
			fmt.Println(bots.CommandErrorText)
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
