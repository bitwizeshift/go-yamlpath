package expr_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/yamlpath/internal/expr"
)

func TestParseSlice(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    *expr.Slice
		wantErr error
	}{
		{
			name:  "Slice does not specify start or end",
			input: ":",
			want: &expr.Slice{
				Start: 0,
				End:   math.MaxInt,
				Step:  1,
			},
		}, {
			name:  "Slice specifies start",
			input: "1:",
			want: &expr.Slice{
				Start: 1,
				End:   math.MaxInt,
				Step:  1,
			},
		}, {
			name:  "Slice specifies end",
			input: ":2",
			want: &expr.Slice{
				Start: 0,
				End:   2,
				Step:  1,
			},
		}, {
			name:  "Slice specifies start and end",
			input: "1:2",
			want: &expr.Slice{
				Start: 1,
				End:   2,
				Step:  1,
			},
		}, {
			name:  "Slice specifies start, end, and step",
			input: "1:2:3",
			want: &expr.Slice{
				Start: 1,
				End:   2,
				Step:  3,
			},
		}, {
			name:    "Invalid slice",
			input:   "hello:world",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := expr.ParseSlice(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseSlice() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseSlice() = %v, want %v", got, want)
			}
		})
	}
}
