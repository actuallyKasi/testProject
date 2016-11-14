package server

import (
	//"fmt"
	"github.com/codegangsta/negroni"
	"metrics/base/server/router"
	//"metrics/base/httputils"
)

func Main()  {
	negroni_instance := negroni.New(negroni.NewRecovery())
	routes := router.SetupRoutes()
	negroni_instance.UseHandler(routes)
	//httputils.SetNegroniInstance(negroni_instance)
	negroni_instance.Run(":" + "8080")
}