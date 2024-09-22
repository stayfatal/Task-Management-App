package endpoints

import (
	"context"
	"tma/services/account/internal/interfaces"
	"tma/services/account/internal/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

type Endpoints struct {
	CreateAccount endpoint.Endpoint
	Login         endpoint.Endpoint
}

func MakeEndpoints(svc interfaces.Service) Endpoints {
	return Endpoints{
		CreateAccount: makeCreateAccountEndpoint(svc),
		Login:         makeLoginEndpoint(svc),
	}
}

func makeCreateAccountEndpoint(svc interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CreateUserRequest)
		token, err := svc.CreateAccount(ctx, req.User)
		if err != nil {
			return models.CreateUserResponse{Err: err}, errors.Wrap(err, "creating account service level")
		}

		return models.CreateUserResponse{Token: token, Err: nil}, nil
	}
}

func makeLoginEndpoint(svc interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.LoginRequest)
		token, err := svc.Login(ctx, req.User)
		if err != nil {
			return models.CreateUserResponse{Err: err}, errors.Wrap(err, "logging service level")
		}

		return models.CreateUserResponse{Token: token, Err: nil}, nil
	}
}
