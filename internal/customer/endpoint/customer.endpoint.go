package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/command"
	"github.com/angryronald/ddd-go-starter/internal/customer/application/query"
	"github.com/angryronald/ddd-go-starter/internal/customer/endpoint/model"
)

type CustomerEndpoint struct {
	customerCommand command.CustomerCommandInterface
	customerQuery   query.CustomerQueryInterface
}

func (e *CustomerEndpoint) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")

	// adding error / http exception handling process
	// ...

	result, _ := e.customerQuery.GetCustomer(r.Context(), uuid.MustParse(customerID))
	resultInBytes, _ := json.Marshal(result)
	w.Write(resultInBytes)

	// adding error handling is a must
	// ...
}

func (e *CustomerEndpoint) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// adding implementation
}

func (e *CustomerEndpoint) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer model.AddingCustomerRequest
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// continue process to parse the endpoint model to application model
	// ...

	// calling customer command functions for addCustomer
	// ...

	// adding http exception handling
	// ...

	// write the result to writer
	// ...
}

func (e *CustomerEndpoint) AddAddress(w http.ResponseWriter, r *http.Request) {
	// adding implementation
}

func (e *CustomerEndpoint) RegisterRoute(r *chi.Mux) *chi.Mux {
	r.Route("/customers", func(r chi.Router) {
		r.Get("/{id}", e.GetCustomer)
		r.Get("/", e.GetCustomers)
		r.Post("/{id}", e.AddCustomer)
		r.Post("/{id}/address", e.AddAddress)
	})
	return r
}

func NewCustomerEndpoint(
	customerCommand command.CustomerCommandInterface,
	customerQuery query.CustomerQueryInterface,
) CustomerEndpointInterface {
	return &CustomerEndpoint{
		customerCommand: customerCommand,
		customerQuery:   customerQuery,
	}
}
