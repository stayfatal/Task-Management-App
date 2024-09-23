package endpoints

import (
	"context"
	"tma/services/task/internal/interfaces"
	"tma/services/task/internal/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

type Endpoints struct {
	AddTask endpoint.Endpoint
}

func MakeEndpoints(svc interfaces.Service) Endpoints {
	return Endpoints{
		AddTask: makeAddTaskEndpoint(svc),
	}
}

func makeAddTaskEndpoint(svc interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CreateTaskRequest)

		err = svc.AddTask(req.Task)

		return models.CreateTaskResponse{Err: err}, errors.Wrap(err, "adding task service level")
	}
}
