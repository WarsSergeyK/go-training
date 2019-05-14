package bots

type Bot interface {
	SayHello() string
	SayTime() string
	SayDate() string
	SayWeekday() string
	SayBye() string
}
