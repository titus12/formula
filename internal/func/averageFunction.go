package _func

import (
	"github.com/xymodule/formula/opt"
	"reflect"
)

type AverageFunction struct {
}

func (*AverageFunction) Name() string {
	return "avg"
}

func (f *AverageFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
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
	if len(args) <= 0{
		return opt.NewArgumentWithType(0, reflect.Float64), nil
	}
	avg := sum / float64(len(args))
	return opt.NewArgumentWithType(avg, reflect.Float64), nil
}
