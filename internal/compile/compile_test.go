package compile_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/yamlpath/internal/compile"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
)

// Note: This test is not evaluating that the expression tree returned from
// compilation is correct, since this is far easier to test in the godog
// cucumber tests. Instead, this test is ensuring that the compilation process
// does not panic or return unexpected errors when compiling valid or invalid
// expressions.

func TestNewTree(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "Invalid root",
			input:   `foo`,
			wantErr: compile.ErrSyntax,
		}, {
			name:  "root expression",
			input: `$`,
		}, {
			name:  "root expression with trailing whitespace",
			input: `$  `,
		}, {
			name:  "field expression",
			input: `$.foo`,
		}, {
			name:  "field wildcard expression",
			input: `$.*`,
		}, {
			name:  "quoted field expression",
			input: `$."foo bar"`,
		}, {
			name:  "recursive descent with field expression",
			input: `$..foo`,
		}, {
			name:  "bracket wildcard expression",
			input: `$[*]`,
		}, {
			name:  "bracket index expression",
			input: `$[0]`,
		}, {
			name:  "bracket slice expression without step",
			input: `$[0:1]`,
		}, {
			name:  "bracket slice expression with step",
			input: `$[0:1:2]`,
		}, {
			name:  "bracket slice expression no start",
			input: `$[:1]`,
		}, {
			name:  "bracket slice expression no end",
			input: `$[0:]`,
		}, {
			name:  "bracket slice expression no start or end",
			input: `$[:]`,
		}, {
			name:  "bracket slice expression no start or end with step",
			input: `$[::2]`,
		}, {
			name:  "bracket union index expression",
			input: `$[0,1]`,
		}, {
			name:  "bracket filter expression with current node",
			input: `$[?(@)]`,
		}, {
			name:  "bracket filter expression with current node and field",
			input: `$[?(@.foo)]`,
		}, {
			name:  "bracket filter expression with == operator",
			input: `$[?(@.foo == 1)]`,
		}, {
			name:  "bracket filter expression with != operator",
			input: `$[?(@.foo != 1)]`,
		}, {
			name:  "bracket filter expression with < operator",
			input: `$[?(@.foo < 1)]`,
		}, {
			name:  "bracket filter expression with <= operator",
			input: `$[?(@.foo <= 1)]`,
		}, {
			name:  "bracket filter expression with > operator",
			input: `$[?(@.foo > 1)]`,
		}, {
			name:  "bracket filter expression with >= operator",
			input: `$[?(@.foo >= 1)]`,
		}, {
			name:  "bracket filter expression with regex",
			input: `$[?(@.foo =~ /bar/)]`,
		}, {
			name:  "bracket filter expression with regex and 'i' flag",
			input: `$[?(@.foo =~ /bar/i)]`,
		}, {
			name:    "bracket filter expression with bad regex",
			input:   `$[?(@.foo =~ /***/`,
			wantErr: compile.ErrSyntax,
		}, {
			name:  "bracket filter expression with boolean literal",
			input: `$[?(true)]`,
		}, {
			name:  "bracket filter expression with integer literal",
			input: `$[?(42 == 42)]`,
		}, {
			name:  "bracket filter expression with float literal",
			input: `$[?(42.5 == 425e-1)]`,
		}, {
			name:  "bracket filter expression with null literal",
			input: `$[?(null)]`,
		}, {
			name:  "bracket filter expression with boolean '&&' operator",
			input: `$[?(true && false)]`,
		}, {
			name:  "bracket filter expression with boolean '||' operator",
			input: `$[?(true || false)]`,
		}, {
			name:  "bracket filter expression with boolean '!' operator",
			input: `$[?(!true)]`,
		}, {
			name:  "bracket filter expression with concat operator '+'",
			input: `$[?(@.foo + @.bar == "hello")]`,
		}, {
			name:  "bracket filter expression with arithmetic operator '-'",
			input: `$[?(@.foo - @.bar == 42)]`,
		}, {
			name:  "bracket filter expression with arithmetic operator '*'",
			input: `$[?(@.foo * @.bar == 42)]`,
		}, {
			name:  "bracket filter expression with arithmetic operator '/'",
			input: `$[?(@.foo / @.bar == 42)]`,
		}, {
			name:  "bracket filter expression with arithmetic operator '%'",
			input: `$[?(@.foo % @.bar == 42)]`,
		}, {
			name:  "bracket filter expression with 'nin' operator",
			input: `$[?(@.foo nin ["bar", "baz"])]`,
		}, {
			name:  "bracket filter expression with 'in' operator",
			input: `$[?(@.foo in ["bar", "baz"])]`,
		}, {
			name:  "bracket filter expression with 'subsetof' operator",
			input: `$[?(@.foo subsetof ["bar", "baz"])]`,
		}, {
			name:  "bracket filter expression with parenthesis",
			input: `$[?(@.foo == 1 && (@.bar == 2 || @.baz == 3))]`,
		}, {
			name:  "bracket filter expression with list literal",
			input: `$[?(@.foo == [1,2,3])]`,
		}, {
			name:  "bracket filter expression with map literal",
			input: `$[?(@.foo == {"bar": 1})]`,
		}, {
			name:  "bracket filter expression with polarity '+' operator",
			input: `$[?(+@.foo == 42)]`,
		}, {
			name:  "bracket filter expression with polarity '-' operator",
			input: `$[?(-@.foo == 42)]`,
		}, {
			name:  "bracket script expression with current node",
			input: `$[(@)]`,
		}, {
			name:  "bracket script expression with field",
			input: `$[(@.foo)]`,
		}, {
			name:  "function expression with correct amount of args",
			input: `$.foo(1,2)`,
		}, {
			name:    "function expression with too few args",
			input:   `$.foo(1)`,
			wantErr: arity.ErrBadArity,
		}, {
			name:    "function expression with too many args",
			input:   `$.foo(1,2,3)`,
			wantErr: arity.ErrBadArity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			table := invocation.NewTable()
			table.Add("foo", nil).SetArity(arity.Exactly(2))
			got, err := compile.NewTree(tc.input, &compile.Config{
				Table: table,
			})

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("NewTree(%q) error = %v, want %v", tc.input, got, want)
			}
			if got, want := (got != nil), (tc.wantErr == nil); got != want {
				t.Errorf("unexpected nil: got %v, want %v", got, want)
			}
		})
	}
}
