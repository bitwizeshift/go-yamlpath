package invocation_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type exampleFunc = func(invocation.Context, ...invocation.Parameter) ([]*yaml.Node, error)

func TestTable_Lookup(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		wantOK bool
	}{
		{
			name:   "Found in base",
			key:    "base",
			wantOK: true,
		}, {
			name:   "Found in derived",
			key:    "derived",
			wantOK: true,
		}, {
			name:   "Not found",
			key:    "not-found",
			wantOK: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			base := invocation.NewTable()
			derived := base.New()
			base.Add("base", nil).SetArity(arity.Any())
			derived.Add("derived", nil).SetArity(arity.Any())

			_, ok := derived.Lookup(tc.key)

			if got, want := ok, tc.wantOK; got != want {
				t.Fatalf("Table.Lookup() ok = %v, want %v", got, want)
			}
		})
	}
}

func TestTable_Invoke(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:   "Calls func",
			params: []invocation.Parameter{nil},
			want:   []*yaml.Node{yamlutil.True},
		}, {
			name:    "Wrong number of params",
			wantErr: arity.ErrBadArity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			table := invocation.NewTable()
			fn := func(invocation.Context, ...invocation.Parameter) ([]*yaml.Node, error) {
				return tc.want, nil
			}
			table.Add("func", fn).SetArity(arity.Exactly(1))

			entry, ok := table.Lookup("func")
			if got, want := ok, true; got != want {
				t.Fatalf("Table.Lookup() unexpected ok: got %v, want %v", got, want)
			}

			got, err := entry.Invoke(nil, tc.params...)
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Entry.Invoke() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Entry.Invoke() = %v, want %v", got, want)
			}
		})
	}
}
