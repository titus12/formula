package fs

import (
	"fmt"
	"github.com/xymodule/formula/opt"
	"reflect"
)

type IfsFunction struct {
}

func (*IfsFunction) Name() string {
	return "ifs"
}

func (f *IfsFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("function %s need more than two arguments", f.Name())
	}

	for i := 0; i < len(args); i+=2 {
		arg, err := (*args[i]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		/*if arg.Type != reflect.Bool {
			return nil, fmt.Errorf("ifs need bool left")
		}*/
		retArg := arg.Bool()
		if retArg.Value.(bool) {
			arg2, _ := (*args[i+1]).Evaluate(context)
			return opt.NewArgument(arg2.Value), nil
		}
	}
	return opt.NewArgumentWithType(0,reflect.Float64), nil
}

func NewIfsFunction() *IfsFunction {
	return &IfsFunction{}
}


