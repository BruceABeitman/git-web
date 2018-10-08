package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

)

// GetRepos retrieves a list of repos from a user using the Git API
func (s ServiceImpl) GetRepos(user string) (*model.Repos, error) {
	response, err := s.WebClient.Get(s.GitAPIURL + "users/" + user + "/repos")
	if err != nil {
		fmt.Printf("Error resolving repos %v", err)
		return nil, ErrResolvingAsset
	}
	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error parsing gitAPI response %v", err)
		return nil, ErrParsingExternalResponse
	}
	defer response.Body.Close()
	repos := new(model.Repos)
	err = json.Unmarshal(jsonResponse, &repos)
	if err != nil {
		fmt.Printf("Error parsing internal repos model %v", err)
		return nil, ErrParsingInternalModel
	}
	return repos, nil
}
