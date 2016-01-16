package api

import (
	"net/http"
)

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello world !"))
}
