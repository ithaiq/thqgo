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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	groups := e.router.groups
	for _, g := range groups {
		for name, handlerMethodMap := range g.handlerMap {
			url := "/" + g.groupName + name
			if r.RequestURI == url {
				ctx := &Context{
					W: w,
					R: r,
				}
				_, ok := handlerMethodMap[ANY]
				if ok {
					handlerMethodMap[ANY](ctx)
					return
				}
				_, ok = handlerMethodMap[r.Method]
				if ok {
					handlerMethodMap[r.Method](ctx)
					return
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (e *Engine) Run() {
	http.Handle("/", e)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
