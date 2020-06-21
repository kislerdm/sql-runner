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
			SELECT col2
			;
			`,
			want: []string{
				"SELECT col1;",
				"SELECT col2;",
			},
		},
		{
			name: "Split Two Queries With Comments",
			in: `SELECT col1; --comment1
			SELECT col2
			;
			`,
			want: []string{
				"SELECT col1;",
				`--comment1
			SELECT col2;`,
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

// func TestStripComments(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		in   string
// 		want string
// 	}{
// 		{
// 			name: "Query With In-line Comments",
// 			in: `select col1, col2 --col3
// from a
// -- don't usually need this
// left join b
// 	using (col1)
// ;`,
// 			want: `select col1, col2
// from a
// left join b
// 	using (col1)
// ;`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		got := sqlparser.StripComments(tt.in)
// 		if got != tt.want {
// 			t.Fatalf("[%s]: Results don't match\ngot: %s\nwant: %s",
// 				tt.name, got, tt.want)
// 		}
// 	}
// }
