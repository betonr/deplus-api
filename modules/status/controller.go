package status

import (
	"net/http"

	ur "github.com/betonr/go-utils/rest"
)

// GetStatus - return status of API
func GetStatus(w http.ResponseWriter, req *http.Request) {
	response := ur.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Running successfully"}
	ur.RespondWithJson(w, http.StatusOK, response)
}
