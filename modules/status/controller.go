package status

import (
	"net/http"

	"github.com/betonr/deplus-api/modules/utils"
)

// GetStatus - return status of API
func GetStatus(w http.ResponseWriter, req *http.Request) {
	response := utils.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Running successfully"}
	utils.RespondWithJson(w, http.StatusOK, response)
}
