package core

import (
	"errors"
	"net/http"

)

// Service blah blah blah
type Service interface {
	GetRepos(string) (*model.Repos, error)
	GetBranches(string, string) ([]model.Branch, error)
	GetSVGs(string, string, string) (*model.SVGList, error)
	GetSVGDetails(string, string, string, string) (*model.SVGDetail, error)
}

// ServiceImpl implements the service interface
type ServiceImpl struct {
	GitAPIURL string
	WebClient *http.Client
}

// NewService constructs a service
func NewService(gitAPIURL string) *ServiceImpl {
	return &ServiceImpl{
		GitAPIURL: gitAPIURL,
		WebClient: http.DefaultClient,
	}
}

var (
	// ErrResolvingAsset an error stating the specified asset does not exist
	ErrResolvingAsset = errors.New("could not resolve requested asset")
	// ErrParsingExternalResponse an error stating the specified repo does not exist
	ErrParsingExternalResponse = errors.New("could not parse an external response")
	// ErrParsingInternalModel an error stating could not parse a model defined in this repo
	ErrParsingInternalModel = errors.New("could not parse internal model")
)

/*
	Provided mocking for this service
*/

// MockServiceImpl ...
type MockServiceImpl struct {
	GitAPIURL                 string
	MockGetBranchesResponse   []model.Branch
	MockGetBranchesErr        error
	MockGetReposResponse      *model.Repos
	MockGetResposErr          error
	MockGetSVGsResponse       *model.SVGList
	MockGetSVGsErr            error
	MockGetSVGDetailsResponse *model.SVGDetail
	MockGetSVGDetailsErr      error
}

// GetRepos ...
func (s MockServiceImpl) GetRepos(user string) (*model.Repos, error) {
	return s.MockGetReposResponse, s.MockGetResposErr
}

// GetBranches ...
func (s MockServiceImpl) GetBranches(user, repo string) ([]model.Branch, error) {
	return s.MockGetBranchesResponse, s.MockGetBranchesErr
}

// GetSVGs ...
func (s MockServiceImpl) GetSVGs(string, string, string) (*model.SVGList, error) {
	return s.MockGetSVGsResponse, s.MockGetSVGsErr
}

// GetSVGDetails ...
func (s MockServiceImpl) GetSVGDetails(string, string, string, string) (*model.SVGDetail, error) {
	return s.MockGetSVGDetailsResponse, s.MockGetSVGDetailsErr
}
