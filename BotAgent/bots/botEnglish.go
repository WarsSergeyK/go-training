package bots

import (
	"time"
)

type BotEnglish struct {
	Name string
}

func (be BotEnglish) SayHello() string {
	return "Hello, I am " + be.Name
}

func (be BotEnglish) SayTime() string {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	return "Now is " + t.Format("15:04:05")
}

func (be BotEnglish) SayDate() string {
	dt := time.Now()
	return "Today is " + dt.Format("January 02, 2006")
}

func (be BotEnglish) SayWeekday() string {
	weekday := time.Now().Weekday().String()
	return "Today is " + weekday
}

func (be BotEnglish) SayBye() string {
	return "Bye"
}

func (br BotEnglish) ParseUserFrase(s string) (string, bool) {
	isExit := false
	response := ""

	switch s {
	case "Hi":
		response = br.SayHello()
	case "Time":
		response = br.SayTime()
	case "Date":
		response = br.SayDate()
	case "Weekday":
		response = br.SayWeekday()
	case "Bye":
		isExit = true
		response = br.SayBye()
	default:
		response = CommandErrorText
	}
	return response, isExit
}
