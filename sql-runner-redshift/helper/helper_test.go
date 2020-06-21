package helper_test

import (
	"fmt"
	"reflect"
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

func TestSQLParametersParser(t *testing.T) {
	tests := []struct {
		in   string
		want map[string]interface{}
		err  error
	}{
		{
			in: `{"a": 1, "b": "col1"}`,
			want: map[string]interface{}{
				"a": 1.,
				"b": "col1",
			},
			err: nil,
		},
		{
			in:   `{"a": 1, "b": "col1"`,
			want: nil,
			err:  fmt.Errorf("unexpected end of JSON input"),
		},
	}

	for _, tt := range tests {
		got, gotErr := helper.SQLParametersParser(tt.in)

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Result outputs don't match\ngot: %v\nwant: %v", got, tt.want)
		}

		switch tt.err == nil {
		case true:
			if gotErr != nil {
				t.Fatalf("Result errors don't match\ngot: %v\nwant: %v", gotErr, tt.err)
			}
		case false:
			if gotErr.Error() != tt.err.Error() {
				t.Fatalf("Result errors don't match\ngot: %v\nwant: %v", gotErr, tt.err)
			}
		}
	}
}
