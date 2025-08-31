package personal_update_status

import (
	"net/http"

	errors "github.com/MediStatTech/MediStat-error"
)

var (
	errInvalidDeparture             = errors.NewHTTPError(http.StatusBadRequest, "Invalid departure.")
	errFailedToUpdatePersonalStatus = errors.NewHTTPError(http.StatusInternalServerError, "Failed to update personal status.")
)
