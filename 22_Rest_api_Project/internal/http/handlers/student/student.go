package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/naman9271/Go-tut/internal/types"
	"github.com/naman9271/Go-tut/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("empty body")))
			return
		}

		slog.Info("New Student created")

		response.WriteJson(w, http.StatusAccepted, map[string]string{"success": "OK"})
	}
}
