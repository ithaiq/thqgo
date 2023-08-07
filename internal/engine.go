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
		for name, handle := range g.handlerMap {
			url := "/" + g.groupName + name
			if r.RequestURI == url {
				ctx := &Context{
					W: w,
					R: r,
				}
				if g.handlerMethodMap[ANY] != nil {
					for _, mName := range g.handlerMethodMap[ANY] {
						if name == mName {
							handle(ctx)
							return
						}
					}
				}
				routers := g.handlerMethodMap[r.Method]
				if routers != nil {
					for _, mName := range routers {
						if name == mName {
							handle(ctx)
							return
						}
					}
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
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
