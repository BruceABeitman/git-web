package api

import (
	"encoding/json"
	"net/http"

	"github.com/goabstract/be-brucebeitman-interview/core"
)

func getSVGs(service core.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user, repo, & branch query param
		user := r.URL.Query().Get("user")
		repo := r.URL.Query().Get("repo")
		branch := r.URL.Query().Get("branch")
		// If there is no user, repo or branch, throw 400
		if len(user) == 0 || len(repo) == 0 || len(branch) == 0 {
			http.Error(w, "user, repo, & branch query params required", http.StatusBadRequest)
			return
		}
		// Get list of branches using user & repo
		blobs, err := service.GetSVGs(user, repo, branch)
		// If internal service errors
		if err != nil {
			handleError(w, err)
			return
		}
		blobsRaw, err := json.Marshal(blobs)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.Write(blobsRaw)
	})
}
