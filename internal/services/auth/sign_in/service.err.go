package sign_in

import (
	"net/http"

	errors "github.com/MediStatTech/MediStat-error"
)

var (
	errInvalidPassword     = errors.NewHTTPError(http.StatusBadRequest, "Invalid password.")
	errInactivePersonal    = errors.NewHTTPError(http.StatusBadRequest, "Personal is inactive.")
	errFailedFindUser      = errors.NewHTTPError(http.StatusInternalServerError, "Failed to find user.")
	errFailedGenerateToken = errors.NewHTTPError(http.StatusInternalServerError, "Failed to generate token.")
)
