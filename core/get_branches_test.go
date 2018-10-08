package core

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

)

type mockWebClient struct {
	http.Client
	mockGetResponse *http.Response
	mockGetError    error
}

func (m mockWebClient) Get(url string) (*http.Response, error) {
	return m.mockGetResponse, m.mockGetError
}

type getBranchesTest struct {
	testName             string
	user                 string
	repo                 string
	mockUpstreamResponse string
	expected             []model.Branch
	expectedError        error
}

var getBranchesTests = []getBranchesTest{
	getBranchesTest{
		testName:             "Valid",
		user:                 "Bruce",
		repo:                 "svg",
		mockUpstreamResponse: okResponse,
		expected: []model.Branch{
			model.Branch{
				Repo: &model.Repo{
					Name: "svg",
				},
				Name:   "develop",
				Commit: "edb346e936663bc9be76ee2cb73ebf3500cb32cf",
			},
		},
		expectedError: nil,
	},
	getBranchesTest{
		testName:             "Invalid - invalid response from git",
		user:                 "Bruce",
		repo:                 "svg",
		mockUpstreamResponse: invalidResponse,
		expected:             nil,
		expectedError:        ErrParsingInternalModel,
	},
}

const (
	okResponse = `[
		{
			"name": "develop",
			"commit": {
				"sha": "edb346e936663bc9be76ee2cb73ebf3500cb32cf",
				"url": "https://api.github.com/repos/BruceABeitman/card-game/commits/edb346e936663bc9be76ee2cb73ebf3500cb32cf"
			}
		}
	]`
	invalidResponse = `{
    "message": "Not Found",
    "documentation_url": "https://developer.github.com/v3/repos/branches/#list-branches"
	}`
)

func TestGetBranches(t *testing.T) {
	for _, tt := range getBranchesTests {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(tt.mockUpstreamResponse))
		})
		mockHTTPClient, teardown, mockURL := testingHTTPClient(h)
		defer teardown()
		// Apparently Golang's http package requires transport in the url?!....ugh
		mockURL = "http://" + mockURL + "/"
		service := &ServiceImpl{
			GitAPIURL: mockURL,
			WebClient: mockHTTPClient,
		}
		actualBranches, actualError := service.GetBranches(tt.user, tt.repo)
		expected, _ := json.Marshal(tt.expected)
		actual, _ := json.Marshal(actualBranches)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\n Failed Response %v\n expected %+v\n actual   %+v", tt.testName, tt.expected, actualBranches)
		}
		if actualError != tt.expectedError {
			t.Errorf("\n Failed Error %v\n expected %+v\n actual %+v", tt.testName, tt.expectedError, actualError)
		}
	}
}

func testingHTTPClient(handler http.Handler) (*http.Client, func(), string) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return cli, s.Close, s.Listener.Addr().String()
}
