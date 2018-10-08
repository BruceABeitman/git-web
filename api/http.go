package api

import (
	"net/http"

	"github.com/goabstract/be-brucebeitman-interview/core"
)

// BuildHTTPServer builds HTTP server
func BuildHTTPServer(service core.Service) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/repos", getRepos(service))
	mux.Handle("/branches", getBranches(service))
	mux.Handle("/svgs", getSVGs(service))
	mux.Handle("/svg/details", getSVGDetails(service))
	return &http.Server{Addr: ":4242", Handler: mux}
}

// Handles responding with the proper HTTP status and response
func handleError(w http.ResponseWriter, err error) {
	switch err {
	case core.ErrResolvingAsset:
		http.Error(w, "Asset was not found", http.StatusBadRequest)
	case core.ErrParsingExternalResponse:
		http.Error(w, "Upstream error", http.StatusBadGateway)
	case core.ErrParsingInternalModel:
		http.Error(w, "Internal error", http.StatusInternalServerError)
	default:
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}
