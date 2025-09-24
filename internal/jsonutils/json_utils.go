package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreposman/action-house-api/internal/validator"
)

func Encode[T any](w http.ResponseWriter, r *http.Request, statusCode int, payload T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("X-Custom-Header", "My custom Header")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("decode json %w", err)
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

func Decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode JSON error: %w", err)
	}

	return data, nil
}
