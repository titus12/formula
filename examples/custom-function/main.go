package main

import (
	"fmt"
	_ "fmt"
	"github.com/titus12/formula"
	"log"
	_ "net/http/pprof"
	"regexp"
	"strings"
	"sync"
	"time"
	//log2 "github.com/sirupsen/logrus"
)

func main() {
	//go http.ListenAndServe("0.0.0.0:6060", nil)
	//log2.Fatal("aaaa")
	/*expression:=formula.NewExpression("[a]+[b]")
	err := expression.AddParameter("a", 1)
	err = expression.AddParameter("b", 2)
	result,err:=expression.Evaluate()
	if err!=nil{
		//handle err
	}

	v,err:= result.Int64()
	if err!=nil{
		//handle err
	}
	log.Printf("result: %v \n",v)
	*/
	//expression:=formula.NewExpression("([x]>[y])?(1+2)*4/2:([z]>[y])?5:6")
	//�ⲿ����
	outerVars := make(map[string]interface{}, 4)
	outerVars["w"] = 21
	outerVars["v"] = 50
	outerVars["z"] = 51
	r, err := HandleFormula("[1]", outerVars)
	fmt.Println(err)
	fmt.Println(r)
	//HandleFormula("[lookup(@x,5,1,100,2,200,3,300,4,400)]",outerVars)
	//HandleFormula("[((2>@xx1)?1:3)]",outerVars)
	//outerVars["x"] = nil
	//HandleFormula("[ifs(@x,1,(max(2,5)>2),2,true,100)]",outerVars)
	//HandleFormula("[1?2:3]",outerVars)
	//HandleFormula("[iif(0&&1,1,2)]",outerVars)
	//HandleFormula("[ifs(a==a,1,2)]",outerVars)
	//outerVars["w"] = "[100-20]" --������
	wg := sync.WaitGroup{}
	count := 1
	wg.Add(count)
	timer := time.NewTicker(3 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			for i := 0; i < count; i++ {
				//func(id int) {
				go func(id int) {
					outerVars := make(map[string]interface{}, 4)
					/*outerVars["win"] = id
					outerVars["match.streak"] = id
					outerVars["1"] = "[@win]"
					outerVars["2"] = "[@match.streak+@win]"
					r, err := HandleFormula("[iif(@match.streak*@1>10,@2,@win)]", outerVars)*/
					//outerVars["match.streak"] = id
					//r, err := HandleFormula("[iif(@match.streak>0,@match.streak,0)]", outerVars)
					/*outerVars["a"] = 3 + id
					outerVars["b"] = 5
					r, err := HandleFormula("[max(lua('"+
						"function max(a,b) "+
						"return math.max(a, b)"+
						" end"+
						"'),"+
						"lua('"+
						"function min(a,b) "+
						"return math.max(a, b)"+
						" end"+
						"')"+
						")]", outerVars)*/
					outerVars["match.win"] = 1
					outerVars["match.lose"] = 1
					outerVars["match.count"] = 10
					outerVars["match.count.quick"] = "[@match.win + 1]"
					outerVars["match.count.ladder"] = 2
					outerVars["match.win.quick"] = "[@match.count + 2]"
					outerVars["match.win.ladder"] = 2
					outerVars["match.lose.aaa.quick"] = 5
					outerVars["match.lose.bbb.quick"] = 2
					outerVars["match.lose.ccc.ladder"] = 3
					//outerVars["match.lose$2$3#2"] = "[count(@match.lose.aaa.quick)]"
					//r, err := HandleFormula("[sum(@match.count.*,@match.win.*,@match.win+2,@match.lose)*2]", outerVars)
					r, err := HandleFormula("[@match.lose*2]", outerVars)
					//r, err := HandleFormula("[avg(@match.win,sum(@match.lose$2$3#2 * 3,10),3,4,5)]", outerVars)
					//wg.Done()
					if err != nil {
						fmt.Printf("%v concurrent error!!! \n", id)
					}
					//if id >= 3 && int(r.(float64)) != outerVars["a"].(int){
					//	fmt.Printf("err ........ err ...... %v \n", id)
					//}

					fmt.Printf("%v [count(@match.win,@match.lose.*.quick,2,3,4,5)] = %v \n", id, r)
				}(i)
			}
			//timer.Stop()
		}
	}
	wg.Wait()

	//outerVars["z"] = "(5*[w])"
	//[lookup(@x,1,100,2,200,3,300....)]
	/*expression:=formula.NewExpression("((2>@xx1)?1:3)")
	//expression:=formula.NewExpression("iif((3>1)||(2>1),1,2)")
	//expression:=formula.NewExpression("iif((3>1),1,2)")
	//expression:=formula.NewExpression("in(2,3,4,5)")
	listener := &CustomParamsListener{}
	err := expression.Precompile(listener)
	if err!=nil{
		//handle err
	}
	params := listener.GetParams()
	for _,p := range params{
		//���ò���
		err = expression.AddParameter(p, outerVars[p])
		if err != nil{
			//handle err
		}
		fmt.Printf("%v \n",p)
	}
	result,err := expression.GetResult()
	v:= result.Value
	if err!=nil{
		//handle err
	}
	log.Printf("result: %v \n",v)*/

	/*v:= result.Value
	if err!=nil{
		//handle err
	}
	log.Printf("result: %v \n",v)*/

	/*var f opt.Function = new(CustomFunction)
	err := formula.Register(&f)
	if err != nil {
		log.Fatal(err)
	}
	outerVars := make(map[string]interface{})
	outerVars["x"] = 1
	expression := formula.NewExpression("CustomFunction([x],2)")
	listener := &CustomParamsListener{}
	err = expression.Precompile(listener)
	if err != nil {
		log.Fatal(err)
	}
	params := listener.GetParams()
	for _,p := range params{
		//���ò���
		err = expression.AddParameter(p, outerVars[p])
		if err != nil{
			//handle err
		}
		fmt.Printf("%v \n",p)
	}
	result,err := expression.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	v, err := result.Int64()
	if err != nil {
		log.Fatal(err)
	}

	if v != 4 { //CustomFunction: i+j+1
		log.Fatal("error")
	}

	log.Println("custom function succeed")

	/*
	var f opt.Function = new(CustomFunction)
	err := formula.Register(&f)
	if err != nil {
		log.Fatal(err)
	}
	expression := formula.NewExpression("CustomFunction(1,2)")
	result, err := expression.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	v, err := result.Int64()
	if err != nil {
		log.Fatal(err)
	}

	if v != 4 { //CustomFunction: i+j+1
		log.Fatal("error")
	}

	log.Println("custom function succeed")

	*/
	log.Println("formula succeed")
}

/*
func Factorial(n int)int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n - 1)
}
*/

func HandleFormula(formulaStr interface{}, vars map[string]interface{}) (result interface{}, err error) {
	defer func() {
		if err := recover(); err != nil { //������panic�쳣
			fmt.Printf("formula err %v\n", err)
		}
	}()
	result, err = ExecFormula(formulaStr, vars)
	return
}

//vars["xx1"] = "[@xx2 * 5]"
//vars["xx2"] = "[3*5]"
//((2>@xx1)?1:3)
func ExecFormula(formulaStr interface{}, vars map[string]interface{}) (result interface{}, err error) {
	if !isFormula(formulaStr) {
		return formulaStr, nil
	}
	newFormulaStr := formulaStr.(string)
	newFormulaStr = newFormulaStr[1 : len(newFormulaStr)-1]
	newFormulaStr = luaFormula(newFormulaStr)
	//newFormulaStr = wildcardParams(newFormulaStr,vars)
	expression := formula.NewExpression(newFormulaStr)
	err = expression.Precompile()
	if err != nil {
		return nil, err
	}
	//��ȡ����
	params := expression.GetParameterNames()
	for _, p := range params {
		if v, ok := vars[p]; ok {
			f, err := ExecFormula(v, vars)
			if err != nil {
				return nil, err
			}
			expression.AddParameter(p, f)
		} else {
			return nil, fmt.Errorf("formula param %v missing", p)
		}
	}
	//��ȡ���
	refResult, err := expression.GetResult()
	if err != nil {
		return nil, err
	}
	v := refResult.Value
	//log.Printf("formula:%v = %v\n",formulaStr,v)
	return v, err
}

func isFormula(formulaStr interface{}) bool {
	if f, ok := formulaStr.(string); ok {
		if strings.HasPrefix(f, "[") {
			return true
		}
	}
	return false
}

func wildcardParams(formulaStr string, vars map[string]interface{}) string {
	var checkWildcard bool
	for _, v := range formula.WildcardSymbols {
		count := strings.Count(formulaStr, v)
		if count > 0 {
			checkWildcard = true
			break
		}
	}
	if !checkWildcard {
		return formulaStr
	}
	exp := regexp.MustCompile(formula.WildcardRegexp)
	matched := exp.FindAllStringSubmatch(formulaStr, -1)
	if len(matched) <= 0 {
		return formulaStr
	}
	for _, v := range matched {
		s := strings.Split(strings.TrimSpace(v[1][1:]), ".")
		var newFormulaStr string
		prefix := strings.TrimSpace(s[0])
	LOOP_PARAMS:
		for k, _ := range vars {
			if strings.HasPrefix(k, prefix) {
				ss := strings.Split(k, ".")
				if len(ss) != len(s) {
					continue LOOP_PARAMS
				}
				for i, sv := range s {
					if sv != ss[i] && sv != "*" {
						continue LOOP_PARAMS
					}
				}
				newFormulaStr += fmt.Sprintf("@%v,", k)
			}
		}
		formulaStr = strings.Replace(formulaStr, v[1], newFormulaStr[0:len(newFormulaStr)-1], 1)
	}
	return formulaStr
}

func luaFormula(formulaStr string) string {
	count := strings.Count(formulaStr, "lua(")
	if count <= 0 {
		return formulaStr
	}
	exp := regexp.MustCompile(formula.LuaFunctionRegexp)
	matched := exp.FindAllStringSubmatch(formulaStr, count)
	if len(matched) > 0 {
		for _, v := range matched {
			newFormulaStr := strings.TrimSpace(v[0][0 : len(v[0])-1])
			params := strings.Split(v[1], ",")
			for _, v := range params {
				newFormulaStr += (",@" + v)
			}
			newFormulaStr += ")"
			formulaStr = strings.Replace(formulaStr, v[0], newFormulaStr, 1)
		}
	}
	return formulaStr
}
