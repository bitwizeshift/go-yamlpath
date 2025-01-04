package funcs

import (
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// paramToInt converts the first argument of the parameter to an integer
func paramToInt(ctx invocation.Context, source string, param invocation.Parameter) (int, error) {
	args, err := param.GetArg(ctx)
	if err != nil {
		return 0, err
	}
	if len(args) != 1 {
		return 0, errs.NewSingletonError(source, args)
	}
	v, err := yamlutil.ToInt(args[0])
	if err != nil {
		return 0, errs.NewEvalError(err)
	}
	return v, nil
}
