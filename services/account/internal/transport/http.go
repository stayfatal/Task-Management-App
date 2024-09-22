package transport

import (
	// "net/http"

	"tma/services/account/internal/endpoints"
	"tma/services/account/internal/interfaces"
	"tma/services/account/internal/middleware"

	gkhttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewServer(svc interfaces.Service) *mux.Router {
	r := mux.NewRouter()

	eps := endpoints.MakeEndpoints(svc)

	r.Methods("POST").Path("/register").Handler(gkhttp.NewServer(
		middleware.DefaultChain()(eps.CreateAccount),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	))

	r.Methods("POST").Path("/login").Handler(gkhttp.NewServer(
		middleware.DefaultChain()(eps.Login),
		decodeLoginRequest,
		encodeLoginResponse,
	))

	return r
}
