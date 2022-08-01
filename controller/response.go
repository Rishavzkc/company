package controller

import (
	"company/model"
	"net/http"
)

var (
	SuccessStatus = model.Response{
		StatusCode: http.StatusOK,
		Message:    "Operation successful",
	}

	MissingMandatoryFields = model.ServiceError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  0001,
		Error:      "missing mandatory fields",
	}

	DBInsertionFailure = model.ServiceError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  1001,
		Error:      "failed to insert in DB",
	}

	DBRetrievalFailure = model.ServiceError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  2001,
		Error:      "failed to retrieve from DB",
	}

	DBUpdateFailure = model.ServiceError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  3001,
		Error:      "failed to update the record in DB",
	}

	DBDeleteFailure = model.ServiceError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  4001,
		Error:      "failed to delete the record in DB",
	}
)
