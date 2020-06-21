package sqlparser

import (
	"regexp"
	"strings"
)

// FormatQuery function to format query
// By replacing placehoders with input config parameters
func FormatQuery(sqlStatement string, parameters map[string]interface{}) (string, error) {

	return string{"a"}, nil
}

// StripComments function to remove comments from queries
func StripComments(input string) string {
	out := regexp.MustCompile(`(--|//)(.*)`).ReplaceAllLiteralString(input, "")
	out = regexp.MustCompile("(?m)^\\s*$[\r\n]*").ReplaceAllLiteralString(out, "")
	return strings.Trim(out, "\r\n")
}

// SplitQueries function to split SQL statements into array of queries
func SplitQueries(sqlStatement string) []string {
	var output []string
	queries := strings.Split(sqlStatement, ";")
	for _, query := range queries[:len(queries)-1] {
		queryOut := strings.TrimSpace(query) + ";"
		output = append(output, queryOut)
	}
	return output
}
