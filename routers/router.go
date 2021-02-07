package routers

import (
	"api-services/controllers"

	"github.com/gorilla/mux"
)

// SetupRouter setups the all router we're going to use into the service
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/getip", controllers.GetIP)
	r.Handle("/api/getip", r)
	return r
}
