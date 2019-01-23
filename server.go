package test_utils

import (
	"net/http"
	"net/http/httptest"
)

// 	HandleFunc the type of func that a typical go http handler has
type HandleFunc func(*http.ResponseWriter, *http.Request)

// 	Routes is a map of functions, that help you to define your testserver
//	Values:
//		key string		The path of your route (e.g "/test/my/cool/func")
//		func HandleFunc	The go http handler that will be served at that route
type Routes map[string]HandleFunc

// 	NewTestServer creates a new go testing server with the given routes, so you can define more complex test server setups in no time
// 	Input
//		routes Routes	The map of routes, that define your test server
func NewTestServer(routes Routes) *httptest.Server {
	routingHandle := func(w http.ResponseWriter, r *http.Request) {
		handleWithRoute(w, r, routes)
	}
	return httptest.NewServer(http.HandlerFunc(routingHandle))
}

func handleWithRoute(w http.ResponseWriter, r *http.Request, routes Routes) {
	path := r.URL.Path
	handle, ok := routes[path]
	if !ok {
		w.WriteHeader(500)
		return
	}
	handle(&w, r)
}
