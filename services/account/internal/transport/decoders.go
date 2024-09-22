package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"tma/services/account/internal/models"

	"github.com/pkg/errors"
)

func decodeCreateUserRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var req models.CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&req.User)
	return req, errors.Wrap(err, "decoding")
}

func decodeLoginRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var req models.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&req.User)
	return req, errors.Wrap(err, "decoding")
}
