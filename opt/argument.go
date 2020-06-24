package opt

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type Argument struct {
	Value interface{}
	Type  reflect.Kind
}

func NewArgument(v interface{}) *Argument {
	if v == nil {
		return &Argument{
			Value: "",
			Type:  reflect.String,
		}
	}

	return &Argument{
		Value: v,
		Type:  reflect.TypeOf(v).Kind(),
	}
}

func NewArgumentWithType(v interface{}, t reflect.Kind) *Argument {
	//todo：此处考虑校验参数类型是否匹配
	//if reflect.TypeOf(v).Kind() != t {
	//}

	return &Argument{
		Value: v,
		Type:  t,
	}
}

func (arg *Argument) Int() (int, error) {
	switch arg.Type {
	case reflect.Int:
		return int(arg.Value.(int)), nil
	case reflect.Int8:
		return int(arg.Value.(int8)), nil
	case reflect.Int16:
		return int(arg.Value.(int16)), nil
	case reflect.Int32:
		return int(arg.Value.(int32)), nil
	case reflect.Int64:
		return int(arg.Value.(int64)), nil
	case reflect.Uint:
		return int(arg.Value.(uint)), nil
	case reflect.Uint8:
		return int(arg.Value.(uint8)), nil
	case reflect.Uint16:
		return int(arg.Value.(uint16)), nil
	case reflect.Uint32:
		return int(arg.Value.(uint32)), nil
	case reflect.Uint64:
		return int(arg.Value.(uint64)), nil
	default:
		return 0, fmt.Errorf("type %s can not convert to int", arg.Type)
	}
}

func (arg *Argument) Int64() (int64, error) {
	switch arg.Type {
	case reflect.Int:
		return int64(arg.Value.(int)), nil
	case reflect.Int8:
		return int64(arg.Value.(int8)), nil
	case reflect.Int16:
		return int64(arg.Value.(int16)), nil
	case reflect.Int32:
		return int64(arg.Value.(int32)), nil
	case reflect.Int64:
		return arg.Value.(int64), nil
	case reflect.Uint:
		return int64(arg.Value.(uint)), nil
	case reflect.Uint8:
		return int64(arg.Value.(uint8)), nil
	case reflect.Uint16:
		return int64(arg.Value.(uint16)), nil
	case reflect.Uint32:
		return int64(arg.Value.(uint32)), nil
	case reflect.Uint64:
		return int64(arg.Value.(uint64)), nil
	case reflect.Float32:
		return int64(arg.Value.(float32)), nil
	case reflect.Float64:
		return int64(arg.Value.(float64)), nil
	default:
		return 0, fmt.Errorf("type %s can not convert to int64", arg.Type)
	}
}

func (arg *Argument) String() string {
	switch arg.Type {
	case reflect.Bool:
		return strconv.FormatBool(arg.Value.(bool))
	case reflect.Int:
		return strconv.Itoa(arg.Value.(int))
	case reflect.Int8:
		return strconv.Itoa(int(arg.Value.(int8)))
	case reflect.Int16:
		return strconv.Itoa(int(arg.Value.(int16)))
	case reflect.Int32:
		return strconv.Itoa(int(arg.Value.(int32)))
	case reflect.Int64:
		return strconv.FormatInt(arg.Value.(int64), 10)
	case reflect.Uint:
		return strconv.Itoa(int(arg.Value.(uint)))
	case reflect.Uint8:
		return strconv.Itoa(int(arg.Value.(uint8)))
	case reflect.Uint16:
		return strconv.Itoa(int(arg.Value.(uint16)))
	case reflect.Uint32:
		return strconv.Itoa(int(arg.Value.(uint32)))
	case reflect.Uint64:
		return strconv.FormatInt(int64(arg.Value.(uint64)), 10)
	case reflect.Float32:
		return strconv.FormatFloat(float64(arg.Value.(float32)), 'g', -1, 64)
	case reflect.Float64:
		return strconv.FormatFloat(arg.Value.(float64), 'g', -1, 64)
	default:
		return fmt.Sprint(arg.Value)
	}
}

func (arg *Argument) Float64() (float64, error) {
	switch arg.Type {
	case reflect.Int:
		return float64(arg.Value.(int)), nil
	case reflect.Int8:
		return float64(arg.Value.(int)), nil
	case reflect.Int16:
		return float64(arg.Value.(int16)), nil
	case reflect.Int32:
		return float64(arg.Value.(int32)), nil
	case reflect.Int64:
		return float64(arg.Value.(int64)), nil
	case reflect.Uint:
		return float64(arg.Value.(uint)), nil
	case reflect.Uint8:
		return float64(arg.Value.(uint8)), nil
	case reflect.Uint16:
		return float64(arg.Value.(uint16)), nil
	case reflect.Uint32:
		return float64(arg.Value.(uint32)), nil
	case reflect.Uint64:
		return float64(arg.Value.(uint64)), nil
	case reflect.Float32:
		return float64(arg.Value.(float32)), nil
	case reflect.Float64:
		return arg.Value.(float64), nil
	default:
		return 0, fmt.Errorf("type %s can not convert to float64", arg.Type)
	}
}

func (arg *Argument) IsNumber() bool {
	return arg.Type == reflect.Int ||
		arg.Type == reflect.Int8 ||
		arg.Type == reflect.Int16 ||
		arg.Type == reflect.Int32 ||
		arg.Type == reflect.Int64 ||
		arg.Type == reflect.Float32 ||
		arg.Type == reflect.Float64 ||
		arg.Type == reflect.Uint ||
		arg.Type == reflect.Uint8 ||
		arg.Type == reflect.Uint16 ||
		arg.Type == reflect.Uint32 ||
		arg.Type == reflect.Uint64
}

func (arg *Argument) IsInteger() bool {
	return arg.Type == reflect.Int ||
		arg.Type == reflect.Int8 ||
		arg.Type == reflect.Int16 ||
		arg.Type == reflect.Int32 ||
		arg.Type == reflect.Int64 ||
		arg.Type == reflect.Uint ||
		arg.Type == reflect.Uint8 ||
		arg.Type == reflect.Uint16 ||
		arg.Type == reflect.Uint32 ||
		arg.Type == reflect.Uint64
}

func (arg *Argument) IsNan() bool {
	v, err := arg.Float64()
	if err != nil {
		return false
	}

	return math.IsNaN(v)
}

func (arg *Argument) Equal(other *Argument) bool {
	if other == nil {
		return false
	}

	if arg.IsNumber() && other.IsNumber() {
		v0, _ := arg.Float64()
		v1, _ := other.Float64()

		return v0 == v1
	}

	if arg.Type == reflect.String && other.Type == reflect.String {
		return arg.String() == other.String()
	}

	return false
}

func (arg *Argument) Bool() *Argument{
	switch arg.Type {
	case reflect.Bool:
		return arg
	case reflect.String:
		if arg.Value == nil || arg.Value == ""{
			return NewArgumentWithType(false,reflect.Bool)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		v,_:= arg.Int()
		if v == 0{
			return NewArgumentWithType(false,reflect.Bool)
		}
	}
	return NewArgumentWithType(true,reflect.Bool)
}

func Float64(value interface{}) (float64, error) {
	switch value.(type) {
	case int,int8:
		return float64(value.(int)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case uint:
		return float64(value.(uint)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	default:
		return 0, fmt.Errorf("type %v can not convert to float64", value)
	}
}