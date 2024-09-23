package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"tma/services/task/internal/models"

	"github.com/pkg/errors"
)

func encodeAddTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(models.CreateTaskResponse)
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	return errors.Wrap(err, "encoding")
}
