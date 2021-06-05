package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/valdirmendesdev/go-hexagonal/adapters/dto"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"net/http"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods(http.MethodPost, http.MethodOptions)

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
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDTO.Name, productDTO.Price)
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}