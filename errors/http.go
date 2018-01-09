package errors

import "net/http"

type NotFound struct {
	httpError
}

func NewNotFound(messages ...string) *NotFound {
	return &NotFound{
		httpError{
			http.StatusNotFound,
			new("Resource not found.", messages...),
		},
	}
}

type BadRequest struct {
	httpError
}

func NewBadRequest(messages ...string) *BadRequest {
	return &BadRequest{
		httpError{
			http.StatusBadRequest,
			new("Invalid request sent.", messages...),
		},
	}
}

type InternalServer struct {
	httpError
}

func NewInternalServer(messages ...string) *InternalServer {
	return &InternalServer{
		httpError{
			http.StatusInternalServerError,
			new("Something was wrong with server.", messages...),
		},
	}
}

type Unauthorized struct {
	httpError
}

func NewUnauthorized(messages ...string) *Unauthorized {
	return &Unauthorized{
		httpError{
			http.StatusUnauthorized,
			new("Unauthorized.", messages...),
		},
	}
}

type Forbidden struct {
	httpError
}

func NewForbidden(messages ...string) *Forbidden {
	return &Forbidden{
		httpError{
			http.StatusForbidden,
			new("Forbidden.", messages...),
		},
	}
}

type UnprocessableEntity struct {
	httpError
}

func NewUnprocessableEntity(messages ...string) *UnprocessableEntity {
	return &UnprocessableEntity{
		httpError{
			http.StatusUnprocessableEntity,
			new("Unprocessable entity.", messages...),
		},
	}
}

type Unavailable struct {
	httpError
}

func NewUnavailable(messages ...string) *Unavailable {
	return &Unavailable{
		httpError{
			http.StatusServiceUnavailable,
			new("Service is currently unavailable.", messages...),
		},
	}
}
