package sqlparser

import (
	"fmt"
	"regexp"
	"strings"
)

// StripComments function to remove comments from queries
func StripComments(input string) string {
	out := regexp.MustCompile(`--(.*)`).ReplaceAllLiteralString(input, "")
	out = regexp.MustCompile("(?m)^\\s*$[\r\n]*").ReplaceAllLiteralString(out, "")
	return strings.Trim(out, "\r\n")
}

// FormatQuery function to format query
// By replacing placehoders with input config parameters
func FormatQuery(sqlStatement string, parameters map[string]interface{}) string {
	args := make([]string, len(parameters)*2)
	i := 0
	for k, v := range parameters {
		args[i] = fmt.Sprintf("{%s}", k)
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(sqlStatement)
}

// SplitQueries function to split SQL statements into array of queries
func SplitQueries(sqlStatement string) []string {
	var output []string
	queries := strings.Split(sqlStatement, ";\n")
	for _, query := range queries {
		query := strings.TrimSpace(query)
		isMatched, _ := regexp.MatchString(";$", query)
		if !isMatched {
			query = fmt.Sprintf("%s;", query)
		}
		output = append(output, query)
	}
	return output
}
