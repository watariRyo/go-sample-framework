package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTables: map[string]*TreeNodes{
				"get":    Constructor(),
				"post":   Constructor(),
				"patch":  Constructor(),
				"put":    Constructor(),
				"delete": Constructor(),
			},
			middlewares: []func(ctx *MyContext){},
		},
	}
}

type Router struct {
	routingTables map[string]*TreeNodes
	middlewares   []func(ctx *MyContext)
	noRoute       func(ctx *MyContext)
}

func (r *Router) Use(middleware func(ctx *MyContext)) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) UseNoRoute(handler func(ctx *MyContext)) {
	r.noRoute = handler
}

func (r *Router) register(method string, pathname string, handler func(ctx *MyContext)) error {
	routingTable := r.routingTables[method]
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := routingTable.Search(pathname)
	if existedHandler != nil {
		panic("already exists handler")
	}
	routingTable.Insert(pathname, handler)
	return nil
}

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	return r.register("get", pathname, handler)
}

func (r *Router) Post(pathname string, handler func(ctx *MyContext)) error {
	return r.register("post", pathname, handler)
}

func (r *Router) Patch(pathname string, handler func(ctx *MyContext)) error {
	return r.register("patch", pathname, handler)
}

func (r *Router) Put(pathname string, handler func(ctx *MyContext)) error {
	return r.register("put", pathname, handler)
}

func (r *Router) Delete(pathname string, handler func(ctx *MyContext)) error {
	return r.register("delete", pathname, handler)
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ctx := NewMyContext(rw, r)

	routingTable := e.Router.routingTables[strings.ToLower(r.Method)]

	pathname := r.URL.Path
	pathname = strings.TrimSuffix(pathname, "/")

	var targetHandler func(ctx *MyContext)
	targetNode := routingTable.Search(pathname)

	if targetNode == nil || targetNode.handler == nil {
		targetHandler = e.Router.noRoute
		if targetHandler == nil {
			defaultNotFoundHandler(ctx)
			return
		}

	} else {
		targetHandler = targetNode.handler
		paramDicts := targetNode.ParaseParams(r.URL.Path)
		ctx.SetParams(paramDicts)
	}

	handlers := append(e.Router.middlewares, targetHandler)
	ctx.SetHandlers(handlers)

	ctx.Next()
}

func (e *Engine) Run() {
	ch := make(chan os.Signal)
	signal.Notify(ch)

	server := &http.Server{Addr: "localhost:8080", Handler: e}

	go func() {
		server.ListenAndServe()
	}()
	<-ch
	fmt.Println("shutdown...")

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("error occurred at shutdown: ", err)
	}
	fmt.Println("shutdown completed")
}

func defaultNotFoundHandler(ctx *MyContext) {
	ctx.ResponseWritor().WriteHeader(http.StatusNotFound)
}
