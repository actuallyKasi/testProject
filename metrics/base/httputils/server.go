package httputils

import "github.com/codegangsta/negroni"

var negroniInstance *negroni.Negroni

func SetNegroniInstance(n *negroni.Negroni) {
	negroniInstance = n
}

