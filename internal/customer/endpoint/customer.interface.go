package endpoint

import (
	"net/http"

	"github.com/go-chi/chi"
)

type CustomerEndpointInterface interface {
	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomers(w http.ResponseWriter, r *http.Request)
	AddCustomer(w http.ResponseWriter, r *http.Request)
	AddAddress(w http.ResponseWriter, r *http.Request)
	RegisterRoute(r *chi.Mux) *chi.Mux
}
