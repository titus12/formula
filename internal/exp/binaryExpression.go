package exp

import (
	"github.com/titus12/formula/internal/cache"
	"github.com/titus12/formula/opt"
)

type BinaryExpression struct {
	Name            string
	LeftExpression  *opt.LogicalExpression
	RightExpression *opt.LogicalExpression
}

func NewBinaryExpression(name string, leftExpression, rightExpression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BinaryExpression{
		Name:            name,
		LeftExpression:  leftExpression,
		RightExpression: rightExpression,
	}

	return &result
}

func (expression *BinaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	f, err := cache.FindFunction(expression.Name)
	if err != nil {
		return nil, err
	}

	return (*f).Evaluate(context, expression.LeftExpression, expression.RightExpression)
}
