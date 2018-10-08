package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

)

var mockService = &core.MockServiceImpl{
	GitAPIURL: "https://api.github.com/",
}

var validBranchesResponse = []model.Branch{
	model.Branch{
		Repo: &model.Repo{
			Name: "svg_test",
		},
		Name:   "develop",
		Commit: "39c914c1031457e19da295d26b31a0e47c7457a6",
	},
	model.Branch{
		Repo: &model.Repo{
			Name: "svg_test",
		},
		Name:   "master",
		Commit: "eca3d3d45ee78cd41ba8c02bf49fe0efe920f9b9",
	},
}
var rawValidBranchesResponse, _ = json.Marshal(validBranchesResponse)

type getBranchesTest struct {
	testName            string
	input               *core.MockServiceImpl
	user                string
	repo                string
	expectedStatus      int
	expectedResponse    string
	expectedError       error
	getBranchesResponse []model.Branch
	getBranchesErr      error
}

var getBranchesTests = []getBranchesTest{
	{
		testName: "Valid",
		input: &core.MockServiceImpl{
			GitAPIURL: "https://api.github.com/",
		},
		user:                "BruceABeitman",
		repo:                "svg_test",
		getBranchesResponse: validBranchesResponse,
		expectedStatus:      http.StatusOK,
		expectedResponse:    string(rawValidBranchesResponse),
		expectedError:       nil,
	},
	{
		testName: "Invalid - Invalid repo query params",
		input: &core.MockServiceImpl{
			GitAPIURL: "https://api.github.com/",
		},
		user:                "BruceABeitman",
		repo:                "",
		getBranchesResponse: nil,
		expectedStatus:      http.StatusBadRequest,
		expectedResponse:    "user & repo query params required",
		expectedError:       nil,
	},
	{
		testName: "Invalid - Invalid user query params",
		input: &core.MockServiceImpl{
			GitAPIURL: "https://api.github.com/",
		},
		user:                "",
		repo:                "svg_t",
		getBranchesResponse: nil,
		expectedStatus:      http.StatusBadRequest,
		expectedResponse:    "user & repo query params required",
		expectedError:       nil,
	},
	{
		testName: "Invalid - internal service error resolving asset",
		input: &core.MockServiceImpl{
			GitAPIURL: "https://api.github.com/",
		},
		user:                "",
		repo:                "svg_t",
		getBranchesResponse: nil,
		getBranchesErr:      core.ErrResolvingAsset,
		expectedStatus:      http.StatusBadRequest,
		expectedResponse:    "Asset was not found",
		expectedError:       nil,
	},
	{
		testName:            "Invalid - internal service error parsing external response",
		input:               mockService,
		user:                "Bruce",
		repo:                "svg_t",
		getBranchesResponse: nil,
		getBranchesErr:      core.ErrParsingExternalResponse,
		expectedStatus:      http.StatusBadGateway,
		expectedResponse:    "Upstream error",
		expectedError:       nil,
	},
	{
		testName:            "Invalid - internal service error parsing internal model",
		input:               mockService,
		user:                "Bruce",
		repo:                "svg_t",
		getBranchesResponse: nil,
		getBranchesErr:      core.ErrParsingInternalModel,
		expectedStatus:      http.StatusInternalServerError,
		expectedResponse:    "Internal error",
		expectedError:       nil,
	},
	{
		testName:            "Invalid - unexpected error",
		input:               mockService,
		user:                "Bruce",
		repo:                "svg_t",
		getBranchesResponse: nil,
		getBranchesErr:      errors.New("Where'd this come from?"),
		expectedStatus:      http.StatusInternalServerError,
		expectedResponse:    "Internal error",
		expectedError:       nil,
	},
}

func TestGetBranches(t *testing.T) {
	for _, tt := range getBranchesTests {
		tt.input.MockGetBranchesErr = tt.getBranchesErr
		tt.input.MockGetBranchesResponse = tt.getBranchesResponse
		handlerFunc := getBranches(tt.input)
		req, err := http.NewRequest("GET", "/branches?user="+tt.user+"&repo="+tt.repo, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerFunc)
		handler.ServeHTTP(rr, req)
		if rr.Code != tt.expectedStatus {
			t.Errorf("\nFailed %v\nHandler returned wrong status code:\n got %v \nwant %v",
				tt.testName,
				rr.Code,
				tt.expectedStatus,
			)
		}
		expected, _ := json.Marshal(tt.expectedResponse)
		if rr.Body.String() == string(expected) {
			t.Errorf(
				"\n Failed %v\n get_branches(%v):\n expected %+v\n actual %+v",
				tt.testName,
				tt.input,
				tt.expectedResponse,
				rr.Body.String(),
			)
		}
	}
}
