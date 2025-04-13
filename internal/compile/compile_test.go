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
			input: `$[0|1]`,
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
			name:  "function with union subexpression",
			input: "$.foo[(1 | 2)]",
		}, {
			name:    "function expression with too few args",
			input:   `$.foo(1)`,
			wantErr: arity.ErrBadArity,
		}, {
			name:    "function expression with too many args",
			input:   `$.foo(1,2,3)`,
			wantErr: arity.ErrBadArity,
		}, {
			name:  "numeric literal",
			input: `1`,
		}, {
			name:  "string literal",
			input: `"foo"`,
		}, {
			name:  "boolean literal",
			input: `true`,
		}, {
			name:  "null literal",
			input: `null`,
		}, {
			name:  "object literal",
			input: `{"foo": "bar"}`,
		}, {
			name:  "array literal",
			input: `["foo", "bar"]`,
		}, {
			name:  "delimited identifier",
			input: `foo.bar-baz`,
		}, {
			name:  "additive expression",
			input: `1 + 2`,
		}, {
			name:  "subtractive expression",
			input: `1 - 2`,
		}, {
			name:  "multiplicative expression",
			input: `1 * 2`,
		}, {
			name:  "divisive expression",
			input: `1 / 2`,
		}, {
			name:  "modulo expression",
			input: `1 % 2`,
		}, {
			name:  "concatenation expression",
			input: `"foo" + "bar"`,
		}, {
			name:  "negative polarity expression",
			input: `-$.foo`,
		}, {
			name:  "positive polarity expression",
			input: `+$.foo`,
		}, {
			name:  "not expression",
			input: `not $.foo`,
		}, {
			name:  "not exclamation expression",
			input: `!$.foo`,
		}, {
			name:  "and expression",
			input: `$.foo and $.bar`,
		}, {
			name:  "or expression",
			input: `$.foo or $.bar`,
		}, {
			name:  "less than expression",
			input: `$.foo < $.bar`,
		}, {
			name:  "less than or equal expression",
			input: `$.foo <= $.bar`,
		}, {
			name:  "greater than expression",
			input: `$.foo > $.bar`,
		}, {
			name:  "greater than or equal expression",
			input: `$.foo >= $.bar`,
		}, {
			name:  "equal expression",
			input: `$.foo == $.bar`,
		}, {
			name:  "not equal expression",
			input: `$.foo != $.bar`,
		}, {
			name:  "in expression",
			input: `$.foo in ["bar", "baz"]`,
		}, {
			name:  "not in expression",
			input: `$.foo nin ["bar", "baz"]`,
		}, {
			name:  "subsetof expression",
			input: `$.foo subsetof ["bar", "baz"]`,
		}, {
			name:  "match expression",
			input: `$.foo =~ /bar/`,
		}, {
			name:  "match expression with flags",
			input: `$.foo =~ /bar/i`,
		}, {
			name:    "match expression with bad regex",
			input:   `$.foo =~ /[a-z/`,
			wantErr: cmpopts.AnyError,
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
