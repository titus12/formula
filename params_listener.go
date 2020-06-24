package formula

import (
	"github.com/titus12/formula/internal/exp"
	"github.com/titus12/formula/internal/parser"
)

type CustomParamsListener struct {
	parser.BaseFormulaListener
	paramNames []string
}

// ExitId is called when exiting the id production.
func (s *CustomParamsListener) ExitId(c *parser.IdContext) {
	v := c.GetRetValue()
	if n, ok := (*v).(*exp.VarIdentifierExpression); ok {
		s.paramNames = append(s.paramNames, n.Name)
	}
}

func (s *CustomParamsListener) GetParamNames() []string {
	return s.paramNames
}

/*func (s *ParamsListener) EnterCalc(c *parser.CalcContext){
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}

// VisitTerminal is called when a terminal node is visited.
func (s *ParamsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *ParamsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *ParamsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *ParamsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}


// EnterExpr is called when entering the expr production.
func (s *ParamsListener)  EnterExpr(c *parser.ExprContext){

}

// EnterOrExpr is called when entering the orExpr production.
func (s *ParamsListener)  EnterOrExpr(c *parser.OrExprContext){

}

// EnterAndExpr is called when entering the andExpr production.
func (s *ParamsListener)  EnterAndExpr(c *parser.AndExprContext){

}

// EnterBitOrExpr is called when entering the bitOrExpr production.
func (s *ParamsListener) EnterBitOrExpr(c *parser.BitOrExprContext){

}

// EnterBitXorExpr is called when entering the bitXorExpr production.
func (s *ParamsListener) EnterBitXorExpr(c *parser.BitXorExprContext){

}

// EnterBitAndExpr is called when entering the bitAndExpr production.
func (s *ParamsListener) EnterBitAndExpr(c *parser.BitAndExprContext){

}

// EnterEqExpr is called when entering the eqExpr production.
func (s *ParamsListener) EnterEqExpr(c *parser.EqExprContext){

}

// EnterRelExpr is called when entering the relExpr production.
func (s *ParamsListener) EnterRelExpr(c *parser.RelExprContext){

}

// EnterShiftExpr is called when entering the shiftExpr production.
func (s *ParamsListener) EnterShiftExpr(c *parser.ShiftExprContext){

}

// EnterAddExpr is called when entering the addExpr production.
func (s *ParamsListener) EnterAddExpr(c *parser.AddExprContext){

}

// EnterMultExpr is called when entering the multExpr production.
func (s *ParamsListener) EnterMultExpr(c *parser.MultExprContext){

}

// EnterUnaryExpr is called when entering the unaryExpr production.
func (s *ParamsListener) EnterUnaryExpr(c *parser.UnaryExprContext){

}

// EnterPrimaryExpr is called when entering the primaryExpr production.
func (s *ParamsListener) EnterPrimaryExpr(c *parser.PrimaryExprContext){

}

// EnterValue is called when entering the value production.
func (s *ParamsListener) EnterValue(c *parser.ValueContext){
	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}

// EnterId is called when entering the id production.
func (s *ParamsListener) EnterId(c *parser.IdContext){
	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}

// EnterΠ is called when entering the π production.
func (s *ParamsListener) EnterΠ(c *parser.ΠContext){

}

// ExitCalc is called when exiting the calc production.
func (s *ParamsListener) ExitCalc(c *parser.CalcContext){
	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}

// ExitExpr is called when exiting the expr production.
func (s *ParamsListener) ExitExpr(c *parser.ExprContext){

}

// ExitOrExpr is called when exiting the orExpr production.
func (s *ParamsListener) ExitOrExpr(c *parser.OrExprContext){

}

// ExitAndExpr is called when exiting the andExpr production.
func (s *ParamsListener) ExitAndExpr(c *parser.AndExprContext){

}

// ExitBitOrExpr is called when exiting the bitOrExpr production.
func (s *ParamsListener) ExitBitOrExpr(c *parser.BitOrExprContext){

}

// ExitBitXorExpr is called when exiting the bitXorExpr production.
func (s *ParamsListener) ExitBitXorExpr(c *parser.BitXorExprContext){

}

// ExitBitAndExpr is called when exiting the bitAndExpr production.
func (s *ParamsListener) ExitBitAndExpr(c *parser.BitAndExprContext){

}

// ExitEqExpr is called when exiting the eqExpr production.
func (s *ParamsListener) ExitEqExpr(c *parser.EqExprContext){

}

// ExitRelExpr is called when exiting the relExpr production.
func (s *ParamsListener) ExitRelExpr(c *parser.RelExprContext){

}

// ExitShiftExpr is called when exiting the shiftExpr production.
func (s *ParamsListener) ExitShiftExpr(c *parser.ShiftExprContext){

}

// ExitAddExpr is called when exiting the addExpr production.
func (s *ParamsListener) ExitAddExpr(c *parser.AddExprContext){

}

// ExitMultExpr is called when exiting the multExpr production.
func (s *ParamsListener) ExitMultExpr(c *parser.MultExprContext){

}

// ExitUnaryExpr is called when exiting the unaryExpr production.
func (s *ParamsListener) ExitUnaryExpr(c *parser.UnaryExprContext){

}

// ExitPrimaryExpr is called when exiting the primaryExpr production.
func (s *ParamsListener) ExitPrimaryExpr(c *parser.PrimaryExprContext){

}

// ExitValue is called when exiting the value production.
func (s *ParamsListener) ExitValue(c *parser.ValueContext){
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}



// ExitΠ is called when exiting the π production.
func (s *ParamsListener) ExitΠ(c *parser.ΠContext){

}*/
