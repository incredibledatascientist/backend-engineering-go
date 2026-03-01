package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ParseIntParam extracts a URL param and converts it to a positive integer.
func ParseIntParam(r *http.Request, key string) (uint, error) {
	vars := mux.Vars(r)

	value, exists := vars[key]
	if !exists || value == "" {
		return 0, fmt.Errorf("missing '%s' parameter", key)
	}

	id64, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id64 == 0 {
		return 0, fmt.Errorf("invalid '%s' parameter: %s", key, value)
	}

	return uint(id64), nil
}

func ParseBody(r *http.Request, v any) error {
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
