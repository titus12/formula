package fs

import (
	"fmt"
	"github.com/xymodule/formula/opt"
	"reflect"
)

type LookupFunction struct {
}

func (*LookupFunction) Name() string {
	return "lookup"
}

func (f *LookupFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("function %s need more than one arguments", f.Name())
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}
	if !arg0.IsNumber() {
		return nil, fmt.Errorf("function %s required numbers", f.Name())
	}
	v0, _ := arg0.Float64()

	for i := 1; i < len(args); i+=2 {
		arg, err := (*args[i]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		if !arg.IsNumber() {
			return nil, fmt.Errorf("function %s required numbers", f.Name())
		}
		v, _ := arg.Float64()
		if v >= v0 {
			arg2, _ := (*args[i+1]).Evaluate(context)
			return opt.NewArgument(arg2.Value), nil
		}
	}
	return opt.NewArgumentWithType(0,reflect.Float64), nil
}

func NewLookupFunction() *LookupFunction {
	return &LookupFunction{}
}


