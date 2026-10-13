package env

import (
	"sync"
	"tp4-debat/com"
)

type Environment struct {
	l sync.Mutex
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
func (env *Environment) Write(message string) {
	env.l.Lock()
	com.StoreConversation(message + "\n")
	env.l.Unlock()
}
