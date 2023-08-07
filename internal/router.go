package internal

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	handlerMap map[string]HandlerFunc
}

func (r *router) Add(name string, handlerFunc HandlerFunc) {
	r.handlerMap[name] = handlerFunc
}
