package utility

import (
	"reflect"
	"testing"

)

type findColorsTest struct {
	testName      string
	input         string
	expected      []model.Color
	expectedError error
}

var findColorsTests = []findColorsTest{
	{
		testName: "rbg%",
		input:    `<svg width="400" height="110"><rect width="300" height="100" style="fill:rgb(200%, 200%, 1%);stroke-width:3;stroke:rgb(0,0,0)" /></svg>`,
		expected: []model.Color{
			model.Color{
				Raw:    "rgb(200%, 200%, 1%)",
				Type:   "rgb",
				Parsed: []string{"200%", "200%", "1%"},
			},
		},
		expectedError: nil,
	},
	{
		testName: "rbg% in a comment",
		input:    `<svg width="400" height="110"><rect width="300" height="100" style="<!-- fill:rgb(200%, 200%, 1%);-->stroke-width:3;stroke:rgb(0,0,0)" /></svg>`,
		expected: []model.Color{
			model.Color{
				Raw:    "rgb(200%, 200%, 1%)",
				Type:   "rgb",
				Parsed: []string{"200%", "200%", "1%"},
			},
		},
		expectedError: nil,
	},
	{
		testName: "named",
		input:    `<svg width="400" height="110"><rect width="300" height="100" style="fill:peru;stroke-width:3;stroke:rgb(0,0,0)" /></svg>`,
		expected: []model.Color{
			model.Color{
				Raw:  "peru",
				Type: "unknown",
			},
		},
		expectedError: nil,
	},
	{
		testName: "rbg int",
		input:    `<svg width="400" height="110"><rect width="300" height="100" style="fill:rgb(0,0,255);stroke-width:3;stroke:rgb(0,0,0)" /></svg>`,
		expected: []model.Color{
			model.Color{
				Raw:    "rgb(0,0,255)",
				Type:   "rgb",
				Parsed: []string{"0", "0", "255"},
			},
		},
		expectedError: nil,
	},
	{
		testName: "icc-color",
		input:    `<circle fill="#CD853F icc-color(acmecmyk, 0.11, 0.48, 0.83, 0.00)"/>`,
		expected: []model.Color{
			model.Color{
				Raw:  "#CD853F icc-color(acmecmyk, 0.11, 0.48, 0.83, 0.00)",
				Type: "unknown",
			},
		},
		expectedError: nil,
	},
	{
		testName: "rgb int 2",
		input:    `<circle fill="rgb(205,133,63)"/>`,
		expected: []model.Color{
			model.Color{
				Raw:    "rgb(205,133,63)",
				Type:   "rgb",
				Parsed: []string{"205", "133", "63"},
			},
		},
		expectedError: nil,
	},
	{
		testName: "Multi: rgb int, named, rgb %, hex",
		input:    `<circle fill="rgb(205,133,63)"/><circle fill="peru"/><circle fill="rgb(80.392%, 52.157%, 24.706%)"/><circle fill="#CD853F"/>`,
		expected: []model.Color{
			model.Color{
				Raw:    "rgb(205,133,63)",
				Type:   "rgb",
				Parsed: []string{"205", "133", "63"},
			},
			model.Color{
				Raw:  "peru",
				Type: "unknown",
			},
			model.Color{
				Raw:    "rgb(80.392%, 52.157%, 24.706%)",
				Type:   "rgb",
				Parsed: []string{"80.392%", "52.157%", "24.706%"},
			},
			model.Color{
				Raw:  "#CD853F",
				Type: "unknown",
			},
		},
		expectedError: nil,
	},
}

func TestFindColors(t *testing.T) {
	for _, tt := range findColorsTests {
		actualColors := FindAllColors(tt.input)
		if !reflect.DeepEqual(actualColors, tt.expected) {
			t.Errorf("\n Failed %v\n FindAllColors(%v):\n expected %+v\n actual   %+v", tt.testName, tt.input, tt.expected, actualColors)
		}
	}
}
