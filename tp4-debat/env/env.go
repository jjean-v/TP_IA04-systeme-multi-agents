package env

import (
	"tp3web"
)

type Environment struct {
}

type Type int

const (
	Question Type = iota
	Answer
)

func (env *Environment) Read() []byte {
	message := tp3web.ReadAll()
	return message
}

func (env *Environment) Write(message string, messageType Type) {
	switch messageType {
	case Question:
		tp3web.StoreConversation("Question: " + message)
	case Answer:
		tp3web.StoreConversation("Answer: " + message)
	}
}
