package api

import (
	"encoding/json"
	"net/http"

)

func getSVGDetails(service core.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user, repo, & branch query param
		user := r.URL.Query().Get("user")
		repo := r.URL.Query().Get("repo")
		branch := r.URL.Query().Get("branch")
		file := r.URL.Query().Get("file")
		// If there is no user, repo or branch, throw 400
		if len(user) == 0 || len(repo) == 0 || len(branch) == 0 || len(file) == 0 {
			http.Error(w, "user, repo, branch, & file query params required", http.StatusBadRequest)
			return
		}
		// Get list of branches using user & repo
		svgDetail, err := service.GetSVGDetails(user, repo, branch, file)
		// If internal service errors
		if err != nil {
			handleError(w, err)
			return
		}
		svgDetailRaw, err := json.Marshal(svgDetail)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.Write(svgDetailRaw)
	})
}
