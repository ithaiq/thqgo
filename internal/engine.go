package internal

import (
	"log"
	"net/http"
)

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router{handlerMap: make(map[string]HandlerFunc)},
	}
}

func (e *Engine) Run() {
	for key, value := range e.handlerMap {
		http.HandleFunc(key, value)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
