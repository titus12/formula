package _func

import (
	"fmt"
	luaf "github.com/xymodule/formula/internal/lua"
	"github.com/xymodule/formula/opt"
	lua "github.com/yuin/gopher-lua"
	"reflect"
	"regexp"
	"strings"
)

type LuaFunction struct {
}

var LuaFunctionNameRegexp = `function([\w\s]+)\(`

func (*LuaFunction) Name() string {
	return "lua"
}

func (f *LuaFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	source := arg0.String()
	code := source[1 : len(source)-1]

	funcName := getFuncName(code)
	if funcName == "" {
		return nil, fmt.Errorf("function %v name is nil..", code)
	}
	//proto cache
	functionProtoCache := luaf.Restore(funcName)
	if functionProtoCache == nil {
		proto, err := luaf.CompileString(code)
		if err != nil {
			return nil, err
		}
		luaf.Store(funcName, proto)
		functionProtoCache = proto
	}
	l := luaf.LuaPool.Get()
	luaf.DoCompiled(l, functionProtoCache)

	/*if err := l.DoString(code); err != nil {
		panic(err)
	}*/

	l.Push(l.GetGlobal(funcName))
	count := len(args)
	for i := 1; i < count; i++ {
		arg1, err := (*args[i]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		if arg1.IsNumber() {
			v, _ := arg1.Float64()
			l.Push(lua.LNumber(v))
		} else {
			v := arg1.String()
			l.Push(lua.LString(v))
		}
	}

	err = l.PCall(count-1, lua.MultRet, nil)
	if err != nil {
		return nil, err
	}
	ret := l.Get(-1)
	l.Pop(1)
	luaf.LuaPool.Put(l)

	if v, ok := ret.(lua.LNumber); ok {
		return opt.NewArgumentWithType(float64(v), reflect.Float64), nil
	} else if v, ok := ret.(lua.LString); ok {
		return opt.NewArgumentWithType(string(v), reflect.String), nil
	}
	return nil, fmt.Errorf("lua result type error neither string nor number")
}

func getFuncName(code string) string {
	exp := regexp.MustCompile(LuaFunctionNameRegexp)
	matched := exp.FindStringSubmatch(code)
	if len(matched) > 1 {
		return strings.TrimSpace(matched[1])
	}
	return ""
}
