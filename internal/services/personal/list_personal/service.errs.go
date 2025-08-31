package list_personal

import (
	"net/http"

	errors "github.com/MediStatTech/MediStat-error"
)

var (
	errFailedToListPersonal = errors.NewHTTPError(http.StatusInternalServerError, "Failed to list personal.")
)
