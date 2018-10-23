package wuery

import (
	"encoding/json"
	"net/http"
)

// HTTPServer implement ServeHTTP interface
type HTTPServer struct {
	w *Wuery
}

// NewHTTPServer return new NewHTTPServer
func NewHTTPServer(w *Wuery) *HTTPServer {
	return &HTTPServer{
		w: w,
	}
}

func (s *HTTPServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	q := Query{}
	if err := json.NewDecoder(req.Body).Decode(&q); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	b, err := s.w.Query(req.Context(), q.Statement)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write(b)
}

// Query DTO
type Query struct {
	Statement string
}
