package formatter

import (
	"html"
	"strings"
)

// JsonStringForCsv cleans and sanitizes a JSON string for CSV.
// It removes unnecessary characters and ensures the JSON string follows valid formatting.
// The function reduces cognitive complexity by utilizing helper functions for specific checks and transformations.
// The cleaned JSON string is returned.
func JsonStringForCsv(jsonString string) (result string) {
	check := strings.Replace(jsonString, " ", "", -1)

	if len(check) < 1 {
		return
	}

	if containsBraces(check) {
		if shouldAddClosingQuote(check) {
			jsonString = jsonString + `\"`
		}
		if shouldRemoveQuoteAndColon(check) {
			jsonString = jsonString[:len(jsonString)-1]
		}
	} else {
		if shouldRemoveFirstQuote(check) {
			jsonString = jsonString[1 : len(jsonString)-0]
		}
	}

	unescapedString := html.UnescapeString(jsonString)
	result = strings.Replace(unescapedString, `,`, ` `, -1)
	result = strings.Replace(result, `;`, ` `, -1)

	return strings.Join(strings.Fields(strings.TrimSpace(result)), " ")
}

// containsBraces checks if the given string contains brace brackets.
// It returns true if braces are present, otherwise false.
func containsBraces(str string) bool {
	return strings.Contains(str, "{") || strings.Contains(str, "}")
}

// shouldAddClosingQuote checks if a closing quote needs to be added to the given string.
// It returns true if a closing quote should be added, otherwise false.
func shouldAddClosingQuote(str string) bool {
	return (str[len(str)-2:] != `"}` && str[len(str)-1:] != `}` && str[len(str)-1:] != `"`) && !strings.Contains(str[len(str)-2:], " ")
}

// shouldRemoveQuoteAndColon checks if the last quote and colon should be removed from the given string.
// It returns true if the quote and colon should be removed, otherwise false.
func shouldRemoveQuoteAndColon(str string) bool {
	return (str[len(str)-2:] != `"}` && str[len(str)-1:] != `}` && str[len(str)-2:] != `\"`) && str[len(str)-2:] == `:"`
}

// shouldRemoveFirstQuote checks if the first quote should be removed from the given string.
// It returns true if the first quote should be removed, otherwise false.
func shouldRemoveFirstQuote(str string) bool {
	return str[0:1] == `"` && str[len(str)-1:] != `"`
}
