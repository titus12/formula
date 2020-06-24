package _func

import (
	"github.com/xymodule/formula/opt"
	"reflect"
)

type SumFunction struct {
}

func (*SumFunction) Name() string {
	return "sum"
}

func (f *SumFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	var sum float64
	for _, arg := range args {
		a, err := (*arg).Evaluate(context)
		if err != nil {
			return nil, err
		}
		v,_ := a.Float64()
		sum += v
	}
	return opt.NewArgumentWithType(sum, reflect.Float64), nil
}