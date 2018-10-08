package api

import (
	"encoding/json"
	"net/http"

)

func getRepos(service core.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user query param
		user := r.URL.Query().Get("user")
		// If there is no user, throw 400
		if len(user) == 0 {
			http.Error(w, "user query param required", http.StatusBadRequest)
			return
		}
		// Get list of repos using user
		repos, err := service.GetRepos(user)
		// If internal service errors
		if err != nil {
			handleError(w, err)
			return
		}
		reposRaw, err := json.Marshal(repos)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.Write(reposRaw)
	})
}
