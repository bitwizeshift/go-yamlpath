package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type FilterExpr struct {
	Expr Expr
}

func (f *FilterExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node
	for _, node := range nodes {
		filtered, err := f.Expr.Eval(ctx, []*yaml.Node{node})
		if err != nil {
			return nil, err
		}

		if len(filtered) > 1 {
			result = append(result, node)
		} else if len(filtered) == 1 {
			b, err := yamlutil.ToBool(filtered[0])
			if err == nil {
				if b {
					result = append(result, node)
				}
			} else {
				// Anything that's not a boolean is considered included
				result = append(result, node)
			}
		}
	}
	return result, nil
}
