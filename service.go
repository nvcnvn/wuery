package wuery

import (
	"net/http"
)

// HTTPServer implement ServeHTTP interface
type HTTPServer struct {
}

func (s *HTTPServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}
