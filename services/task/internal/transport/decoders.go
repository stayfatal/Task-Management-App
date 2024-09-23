package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"tma/services/task/internal/models"

	"github.com/pkg/errors"
)

func decodeAddTaskRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var req models.CreateTaskRequest
	err := json.NewDecoder(request.Body).Decode(&req.Task)
	return req, errors.Wrap(err, "decoding")
}
