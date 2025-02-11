package requests

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ValidateAdminId(r *http.Request) (uint, error) {
	params := mux.Vars(r)
	adminID := params["id"]

	id, err := strconv.ParseUint(adminID, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
