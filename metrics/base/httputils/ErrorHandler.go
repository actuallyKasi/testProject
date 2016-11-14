package httputils

type HttpErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}