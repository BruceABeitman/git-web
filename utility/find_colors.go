package utility

import (
	"regexp"
	"strings"

)

// Regex for matching against 'fill', followed by any single non-alphanumeric, followed by anything until any of the following: ;/<>
var findFill = regexp.MustCompile("fill[^a-zA-Z\\d][^;/<>]*")

// Regex for matching against parentheticals
var findParenth = regexp.MustCompile(`\((.*?)\)`)

// FindAllColors given the contents of a file, returns all color objects within that file
// (as delineated from findFill regex)
func FindAllColors(contents string) []model.Color {
	matches := findFill.FindAllString(contents, -1)
	colors := make([]model.Color, len(matches))
	for index, match := range matches {
		colors[index] = buildColor(match)
	}
	return colors
}

// Given a raw color string, build the parsed color object
func buildColor(colorString string) model.Color {
	// Pull off the first 5 chars (fill=, fill:, etc.)
	// And remove any double quotes
	raw := strings.Trim(colorString[5:], "\"")
	colorType := "unknown"
	parsed := make([]string, 0)
	// If color is RGB type, parse RGB type
	if strings.Contains(raw, "rgb(") {
		colorType = "rgb"
		// Resolve the components from the parenthetical
		details := findParenth.FindString(raw)
		// Remove parenthesis
		details = strings.Trim(details, "()")
		// Split out color components "(a, b, c)" => [a, b, c]
		parsed = strings.Split(details, ",")
		// Remove all spaces from each parsed element
		for index := range parsed {
			parsed[index] = strings.TrimSpace(parsed[index])
		}
	}
	color := model.Color{
		Raw:  raw,
		Type: colorType,
	}
	// If we parsed the color components, populate it on the response
	if len(parsed) > 0 {
		color.Parsed = parsed
	}
	return color
}
