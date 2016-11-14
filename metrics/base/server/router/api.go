package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

const(
	BasePath = "/metrics"
)
var computeRouters map[string]*mux.Router
var rootRouter *mux.Router
//var baseRouter *mux.Router

type EntityRouter interface {
	RegisterRoutes()
}

type Router struct {
	entityRouters []EntityRouter
}

func NewRouter(er ...EntityRouter) *Router {
	return &Router{entityRouters: er}
}


func (r *Router) Register() {
	for _, er := range r.entityRouters {
		er.RegisterRoutes()
	}
}

func RootRouter() *mux.Router {
	return rootRouter
}

func createRouters(){
	rootRouter = mux.NewRouter().StrictSlash(true)
	//baseRouter = rootRouter.PathPrefix(BasePath).Subrouter().StrictSlash(true)

}
func SetupRoutes() http.Handler{
	createRouters()
	NewRouter(
		NewSystemMetricsRouter(),
	).Register()
	http.Handle("/", RootRouter())
	return http.DefaultServeMux
}