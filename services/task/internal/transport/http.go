package transport

import (
	"tma/services/task/internal/endpoints"
	"tma/services/task/internal/interfaces"
	"tma/services/task/internal/middleware"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewServer(svc interfaces.Service) *mux.Router {
	r := mux.NewRouter()

	eps := endpoints.MakeEndpoints(svc)

	r.Methods("POST").Path("/task").Handler(kithttp.NewServer(
		middleware.DefaultChain()(eps.AddTask),
		decodeAddTaskRequest,
		encodeAddTaskResponse,
	))

	return r
}
