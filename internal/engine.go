package internal

import (
	"log"
	"net/http"
)

type Engine struct {
	*router
}

func New() *Engine {
	return &Engine{
		newRouter(),
	}
}

func (e *Engine) Run() {
	groups := e.router.groups
	for _, g := range groups {
		for key, value := range g.handlerMap {
			http.HandleFunc("/"+g.groupName+key, value)
		}
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
