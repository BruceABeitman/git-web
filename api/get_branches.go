package api

import (
	"encoding/json"
	"net/http"

)

func getBranches(service core.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user & repo query param
		user := r.URL.Query().Get("user")
		repo := r.URL.Query().Get("repo")
		// If there is no user or repo, throw 400
		if len(user) == 0 || len(repo) == 0 {
			http.Error(w, "user & repo query params required", http.StatusBadRequest)
			return
		}
		// Get list of branches using user & repo
		branches, err := service.GetBranches(user, repo)
		// If internal service errors
		if err != nil {
			handleError(w, err)
			return
		}
		branchesRaw, err := json.Marshal(branches)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.Write(branchesRaw)
	})
}
