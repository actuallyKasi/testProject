package httputils

import (
	"net/http"
	//"metrics/base/utils"
)

func WriteResponse(writer *http.ResponseWriter, encoding string, response []byte, ReturnCode int) {
	(*writer).Header().Set("Content-Type", "application/"+encoding)
	if ReturnCode >= 200 && ReturnCode < 300 {
		(*writer).Write(response)
		return
	}
	http.Error(*writer, string(response), http.StatusBadRequest)
	return
}



