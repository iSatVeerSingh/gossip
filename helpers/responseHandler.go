package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseType struct {
	Success bool        `json:"success"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

func GetErrorResponse(w http.ResponseWriter, data interface{}, code int) {
	newError := ResponseType{
		Success: false,
		Error:   true,
		Data:    data,
	}

	result, err := json.Marshal(newError)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(result)
}
func GetSuccessResponse(w http.ResponseWriter, data interface{}, code int) {
	newSuccess := ResponseType{
		Success: true,
		Error:   false,
		Data:    data,
	}

	result, err := json.Marshal(newSuccess)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(result)
}
