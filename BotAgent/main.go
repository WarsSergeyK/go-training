package main

import (
	"botagent/domain"
	"fmt"
)

func main() {

	fmt.Print("Language: ")

	var language string
	_, err := fmt.Scanln(&language)
	if err != nil {
		fmt.Println("Alarma!!!")
	}

	bot := createBot(language)
	if bot == nil {
		fmt.Println("Unknown language!!!")
		return
	}

	answer := 0
	for {
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Alarma!!!")
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
			fmt.Println("Unknown command")
		}
	}
}

func createBot(language string) domain.Bot {

	var b domain.Bot

	switch language {
	case "English":
		b = domain.BotEnglish{Name: "John"}
	case "Russian":
		b = domain.BotRussian{Name: "Иван"}
		// default:
		// "Unknown language"
	}
	return b
}
