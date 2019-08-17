package status

import (
	"net/http"

	r "github.com/betonr/go-utils/rest"
)

// GetStatus - return status of API
func GetStatus(w http.ResponseWriter, req *http.Request) {
	response := r.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Running successfully"}
	r.RespondWithJson(w, http.StatusOK, response)
}
