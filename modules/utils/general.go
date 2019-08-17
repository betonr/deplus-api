package utils

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// GetBetween - return the first item different of empty
func GetBetween(items []string) string {
	for _, item := range items {
		if item != "" {
			return item
		}
	}
	return ""
}

// ContainStr - return true if s contain e
func ContainStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// RespondWithJson - response formated with json
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Validate - convert/get payload by json
func Validate(w http.ResponseWriter, object interface{}) bool {
	result, err := govalidator.ValidateStruct(object)
	if !result && err != nil {
		response := DefaultResponse{http.StatusBadRequest, err.Error()}
		RespondWithJson(w, http.StatusBadRequest, response)
		return false
	}
	return true
}
