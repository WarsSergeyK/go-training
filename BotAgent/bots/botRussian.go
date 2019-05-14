package bots

import (
	"time"
)

type BotRussian struct {
	Name string
}

func (br BotRussian) SayHello() string {
	return "Привет, я " + br.Name
}

func (br BotRussian) SayTime() string {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	return "Сейчас " + t.Format("15:04:05")
}

func (br BotRussian) SayDate() string {
	dt := time.Now()
	return "Сегодня " + dt.Format("02 January 2006") + " года"
}

func (br BotRussian) SayWeekday() string {
	weekday := time.Now().Weekday().String()
	return "Сегодня " + weekday
}

func (br BotRussian) SayBye() string {
	return "Пока"
}
