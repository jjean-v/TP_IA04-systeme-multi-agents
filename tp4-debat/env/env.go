package env

import (
	"tp4-debat/com"
)

type Environment struct {
}

type Type int

const (
	Question Type = iota
	Answer
)

var test = Answer

func NewEnvironment() *Environment {
	return &Environment{}
}

func (env *Environment) Read() []byte {
	message := com.ReadAll()
	return message
}

// Need to add the Mutex
func (env *Environment) Write(message string, messageType Type) {
	switch messageType {
	case Question:
		com.StoreConversation("Message " + message)
	case Answer:
		com.StoreConversation("Answer " + message)
	}
}
