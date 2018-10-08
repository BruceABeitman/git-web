package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

)

// GetBranches retrieves a list of branches, from Git API, given a user and repo
func (s ServiceImpl) GetBranches(user, repo string) ([]model.Branch, error) {
	response, err := s.WebClient.Get(s.GitAPIURL + "repos/" + user + "/" + repo + "/branches")
	if err != nil {
		fmt.Printf("Error resolving branches %v", err)
		return nil, ErrResolvingAsset
	}
	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error parsing gitAPI response %v", err)
		return nil, ErrParsingExternalResponse
	}
	defer response.Body.Close()
	gitBranches := make([]model.GitBranch, 0)
	err = json.Unmarshal(jsonResponse, &gitBranches)
	if err != nil {
		fmt.Printf("Error parsing internal gitBranches model %v", err)
		return nil, ErrParsingInternalModel
	}
	branches := make([]model.Branch, len(gitBranches))
	for index, branch := range gitBranches {
		branches[index] = model.Branch{
			Repo: &model.Repo{
				Name: repo,
			},
			Name:   branch.Name,
			Commit: branch.Commit.Sha,
		}
	}
	return branches, nil
}
