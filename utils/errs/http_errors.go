package errs

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	BadRequestStr          = "Bad request"
	NotFoundStr            = "Not Found"
	BadQueryParamsStr      = "Invalid query params"
	InternalServerErrorStr = "Internal Server Error"
	StatusNotModified      = "Status Not Modified"
	StatusConflict         = "Status Conflict"
	InvalidScopes          = "Invalid scope"
)

var (
	ErrMethodNotAllowed  = &ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrNotFound          = &ErrorResponse{StatusCode: http.StatusNotFound, Message: "Resource not found"}
	ErrBadRequest        = &ErrorResponse{StatusCode: http.StatusBadRequest, Message: BadRequestStr}
	ErrBadQueryParams    = &ErrorResponse{StatusCode: http.StatusBadRequest, Message: BadQueryParamsStr}
	ErrStatusNotModified = &ErrorResponse{StatusCode: http.StatusNotModified, Message: StatusNotModified}
	ErrStatusConflict    = &ErrorResponse{StatusCode: http.StatusConflict, Message: StatusNotModified}
	ErrInvalidScopesStr  = &ErrorResponse{StatusCode: http.StatusUnauthorized, Message: InvalidScopes}
)

// error response
type ErrorResponse struct {
	Err        error  `json:"-"`                     // low-level runtime error
	StatusCode int    `json:"-"`                     // http response status code
	StatusText string `json:"status_text,omitempty"` // http response status text
	Message    string `json:"message"`               // application-level error message
}

// region public mehtods
// override render
// @param w http.ResponseWriter
// @param r *http.Request
// @result error
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
} // Austin 20220202

// new bad request error renderer
// @param err error
func BadRequestErrRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusBadRequest,
		StatusText: BadRequestStr,
		Message:    err.Error(),
	}
} // Austin 20220202

// new not found error renderer
// @param err error
func NotFoundErrRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusNotFound,
		StatusText: NotFoundStr,
		Message:    err.Error(),
	}
} // Austin 20220202

// server erroer renderer
// @param err error
func ServerErrRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
		StatusText: InternalServerErrorStr,
		Message:    err.Error(),
	}
} // Austin 20220202

// not modified erroer renderer
// @param err error
func NotModifiedErrRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusNotModified,
		StatusText: StatusNotModified,
		Message:    err.Error(),
	}
} // Eden 20220418

// conflict erroer renderer
// @param err error
func ConflictErrRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusConflict,
		StatusText: StatusConflict,
		Message:    err.Error(),
	}
} // Eden 20220418

// endregion
