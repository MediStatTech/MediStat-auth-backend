package get_personal

import (
	"net/http"

	errors "github.com/MediStatTech/MediStat-error"
)

var (
	errInactivePersonal = errors.NewHTTPError(http.StatusBadRequest, "Personal is inactive.")
	errFailedToFindUser = errors.NewHTTPError(http.StatusInternalServerError, "Failed to find user.")
)
