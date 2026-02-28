package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ParseIntParam extracts a URL param and converts it to a positive integer.
func ParseIntParam(r *http.Request, key string) (int, error) {
	vars := mux.Vars(r)

	value, exists := vars[key]
	if !exists || value == "" {
		return 0, fmt.Errorf("missing '%s' parameter", key)
	}

	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid '%s' parameter: %s", key, value)
	}

	return id, nil
}
