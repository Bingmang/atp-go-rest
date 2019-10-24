package routes

import (
	"atp-go-rest/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method string
	Pattern string
	Handler http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes[] Route

func init() {
	register("GET", "/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Hi!")
	}, nil)
	register("POST", "/coach", controllers.AddCoach, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
