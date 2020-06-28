package sqlparser_test

import (
	"testing"

	"github.com/kislerdm/sql-runner/sql-runner-redshift/sqlparser"
)

func TestSplitQueries(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "Split Two SQL Queries",
			in: `SELECT col1;
			SELECT col2;`,
			want: []string{
				"SELECT col1;",
				"SELECT col2;",
			},
		},
	}

	for _, tt := range tests {
		got := sqlparser.SplitQueries(tt.in)
		if len(tt.want) != len(got) {
			t.Fatalf("[%s]: Dims don't match", tt.name)
		} else {
			for i := 0; i < len(tt.want); i++ {
				if tt.want[i] != got[i] {
					t.Fatalf("[%s]: Results don't match\ngot: %s\nwant: %s",
						tt.name, got[i], tt.want[i])
				}
			}
		}
	}
}

func FormatQuery(t *testing.T) {
	tests := []struct {
		inQuery      string
		inParameters map[string]interface{}
		want         string
	}{
		{
			inQuery: `SELECT {col1} as a, {col2} as b, {col3} as date FROM data;`,
			inParameters: map[string]interface{}{
				"col1": "date",
				"col2": 1,
				"col3": "current_date",
			},
			want: `SELECT date as a, 1 as b, current_date as date FROM data;`,
		},
	}
	for _, tt := range tests {
		got := sqlparser.FormatQuery(tt.inQuery, tt.inParameters)
		if tt.want != got {
			t.Fatalf("Results don't match\ngot: %s\nwant: %s",
				got, tt.want)
		}
	}
}

// func TestStripComments(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		in   string
// 		want string
// 	}{
// 		{
// 			name: "Query With In-line Comments",
// 			in: `select col1, col2
// from a -- test
// ;`,
// 			want: `select col1, col2
// from a
// ;`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		got := sqlparser.StripComments(tt.in)
// 		if !strings.EqualFold(got, tt.want) {
// 			t.Fatalf("[%s]: Results don't match\ngot: %s\nwant: %s",
// 				tt.name, got, tt.want)
// 		}
// 	}
// }
