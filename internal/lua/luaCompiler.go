package lua

import (
	"bufio"
	"errors"
	glua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"strings"
	"sync"
)

var _default_function_proto FunctionProtoManager

type FunctionProtoManager struct {
	functionProto map[string]*glua.FunctionProto
	mu sync.RWMutex
}

func init(){
	_default_function_proto.init()
}

func (f *FunctionProtoManager) init(){
	f.functionProto = make(map[string]*glua.FunctionProto,255)
}

func (f *FunctionProtoManager) add(key string, value *glua.FunctionProto){
	defer f.mu.Unlock()
	f.mu.Lock()
	f.functionProto[key] = value
}

func (f *FunctionProtoManager) get(key string) *glua.FunctionProto{
	defer f.mu.RUnlock()
	f.mu.RLock()
	return f.functionProto[key]
}

func Store(key string, functionProto *glua.FunctionProto) {
	_default_function_proto.add(key,functionProto)
}

func Restore(key string) *glua.FunctionProto{
	return _default_function_proto.get(key)
}


//compile lua string
func CompileString(source string) (*glua.FunctionProto, error) {
	reader := strings.NewReader(source)
	chunk, err := parse.Parse(reader, source)
	if err != nil {
		return nil, err
	}
	proto, err := glua.Compile(chunk, source)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

func DoCompiled(L *glua.LState, proto *glua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, glua.MultRet, nil)
}

//compile lua file
func CompileFile(filePath string) (*glua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := glua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// 获取对应指令的 OpCode
func opGetOpCode(inst uint32) int {
	return int(inst >> 26)
}

func CheckGlobal(proto *glua.FunctionProto) error {
	for _, code := range proto.Code {
		switch opGetOpCode(code) {
		case glua.OP_GETGLOBAL:
			return errors.New("not allow to access global")
		case glua.OP_SETGLOBAL:
			return errors.New("not allow to set global")
		}
	}
	// 对嵌套函数进行全局变量的检查
	for _, nestedProto := range proto.FunctionPrototypes {
		if err := CheckGlobal(nestedProto); err != nil {
			return err
		}
	}
	return nil
}

