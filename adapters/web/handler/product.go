package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"golang.org/x/text/encoding"
	"net/http"
)

func NewProductHandler(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface)  {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
		)).Methods(http.MethodGet, http.MethodOptions)
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		
	})
}