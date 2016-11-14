package errors

import (
	"text/template"
	"bytes"
)
// The Error interface type is our custom interface for representing
// an error condition, with the nil value representing no error.
type Error interface {
	error
	LogMessage() string
}

// CustomError is our error definition.
type CustomError struct {
	Message string
	Cause   string
}
// Error implements the error interface.
func (e *CustomError) Error() string {
	return e.Message
}
var logErrorTemplate = template.Must(template.New("logErrorTemplate").Parse(
	`Error: {{.Message}}, Cause: {{.Cause}}`,
))

// LogMessage implements the custom Error interface
func (e *CustomError) LogMessage() string {
	var logBuffer bytes.Buffer
	logErrorTemplate.Execute(&logBuffer, e)
	return logBuffer.String()
}

func New(message string, cause string) *CustomError {
	return &CustomError{Message: message, Cause: cause}
}
