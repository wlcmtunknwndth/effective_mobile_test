package httpResp

import "net/http"

func WriteResponse(w http.ResponseWriter, statusCode int, info string) {
	_, err := w.Write([]byte(info))
	if err != nil {
		return
	}
	w.WriteHeader(statusCode)
}
