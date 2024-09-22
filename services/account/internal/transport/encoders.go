package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"tma/services/account/internal/models"

	"github.com/pkg/errors"
)

func encodeCreateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(models.CreateUserResponse)
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(resp)

	return errors.Wrap(err, "encoding")
}

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(models.LoginResponse)
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(resp)

	return errors.Wrap(err, "encoding")
}
