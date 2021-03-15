package http_service

import "net/http"

func createHTTPServer(svc *HTTPServiceAPI) *http.Server {
	var svr http.Server

	mux := http.NewServeMux()

	mux.HandleFunc("/version", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.WriteHeader(http.StatusOK)
			res.Header().Add("Content-Type", "text-plain")
			res.Write([]byte(svc.Version))
		} else {
			res.WriteHeader(http.StatusMethodNotAllowed)
			res.Header().Add("Content-Type", "text-plain")
			res.Write([]byte("Only 'GET' requests are supported"))
		}
	})

	svr.Handler = mux

	return &svr
}
