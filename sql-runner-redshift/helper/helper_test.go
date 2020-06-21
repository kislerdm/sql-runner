package helper_test

import (
	"testing"

	"github.com/kislerdm/sql-runner/sql-runner-redshift/helper"
)

func TestInArrayStr(t *testing.T) {
	tests := []struct {
		array       []string
		testElement string
		want        bool
	}{
		{
			array:       []string{"a", "b", "asd"},
			testElement: "b",
			want:        true,
		},
		{
			array:       []string{"a", "b", "asd"},
			testElement: "check",
			want:        false,
		},
	}

	for _, tt := range tests {
		got := helper.InArrayStr(tt.array, tt.testElement)
		if tt.want != got {
			t.Fatalf("Results don't match\ngot: %v\nwant: %v", got, tt.want)
		}
	}
}
