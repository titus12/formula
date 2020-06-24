package _func

import (
"github.com/xymodule/formula/opt"
"reflect"
)

type CountFunction struct {
}

func (*CountFunction) Name() string {
	return "count"
}

func (f *CountFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}
	count := len(args)
	return opt.NewArgumentWithType(count, reflect.Int), nil
}