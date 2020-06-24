package fs

import (
	"github.com/xymodule/formula/opt"
	"reflect"
)

type NeFunction struct {
}

func (*NeFunction) Name() string {
	return "!="
}

func (f *NeFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)

	if !arg0.Equal(arg1) {
		return opt.NewArgumentWithType(true, reflect.Bool), nil
	}
	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewNeFunction() *NeFunction {
	return &NeFunction{}
}
