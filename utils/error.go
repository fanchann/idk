package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"fanchann/elegant_err/models"
)

func ErrorAPIHandler[T any](value T, err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("error : %v\n", err)
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// success response
	response := models.APIResponse{
		Success: true,
		Data:    value,
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}

func CatchError[T any](value T, err error) T {
	if err != nil {
		log.Printf("error : %v\n", err)
		var zeroValue T
		return zeroValue
	}
	return value
}
