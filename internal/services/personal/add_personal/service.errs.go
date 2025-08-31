package add_personal

import (
	"net/http"

	errors "github.com/MediStatTech/MediStat-error"
)

var (
	errUserAlreadyExists      = errors.NewHTTPError(http.StatusConflict, "User already exists.")
	errInvalidDeparture       = errors.NewHTTPError(http.StatusBadRequest, "Invalid departure.")
	errFailedToFindUser       = errors.NewHTTPError(http.StatusInternalServerError, "Failed to find user.")
	errFailedToHashPassword   = errors.NewHTTPError(http.StatusInternalServerError, "Failed to hash password.")
	errFailedToCreatePersonal = errors.NewHTTPError(http.StatusInternalServerError, "Failed to create personal.")
)
