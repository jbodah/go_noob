package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type SqlLogger struct {
	Name string
}

func (l SqlLogger) Log(message string) {
	fmt.Println("[SqlLogger]", message)
}

type ConsoleLogger struct {
	Name string
}

func (l ConsoleLogger) Log(message string) {
	fmt.Println("[ConsoleLogger]", message)
}

type FileLogger struct {
	Name string
}

func (l FileLogger) Log(message string) {
	fmt.Println("[FileLogger]", message)
}

type Server struct {
	Logger Logger
}

func process(logger Logger) {
	logger.Log("hello!")
}

func main() {
	s := Server{FileLogger{"FileLogger"}}
	process(s.Logger)

	s = Server{ConsoleLogger{"ConsoleLogger"}}
	process(s.Logger)

	s = Server{SqlLogger{"SqlLogger"}}
	process(s.Logger)
}
