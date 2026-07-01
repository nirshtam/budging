package http

import "net/http"

func startHttpServer() {
	apiMux := http.NewServeMux()

	http.ListenAndServe(":8000", apiMux)
}
