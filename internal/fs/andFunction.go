package fs

import (
	"fmt"
	"github.com/xymodule/formula/opt"
	"reflect"
)

type AndFunction struct {
}

func (*AndFunction) Name() string {
	return "&&"
}

func (f *AndFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("function %s required three arguments", f.Name())
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	/*if arg0.Type != reflect.Bool {
		return nil, fmt.Errorf("the first argument of function %s should be bool", f.Name())
	}*/
	retArg := arg0.Bool()
	if !retArg.Value.(bool) {
		return opt.NewArgumentWithType(false, reflect.Bool), nil
	}
	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	/*if arg1.Type != reflect.Bool {
		return nil, fmt.Errorf("the second argument of function %s should be bool", f.Name())
	}*/
	retArg = arg1.Bool()
	if !retArg.Value.(bool) {
		return opt.NewArgumentWithType(false, reflect.Bool), nil
	}
	return opt.NewArgumentWithType(true, reflect.Bool), nil
}

func NewAndFunction() *AndFunction {
	return &AndFunction{}
}
