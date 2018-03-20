package expressions

import (
	"fmt"
	"github.com/ckLXHL/gotoy/src/token"
)

//expression格式化接口
type Expression interface {
	fmt.Stringer
}

//基础表达式
type NameExpression struct {
	name string
}

func (ne NameExpression) String() string {
	return ne.name
}
/*
//条件表达式
type ConditionalExpression struct{
	condition, thenExp, elseExp ExpressionIf
}

func (ce ConditionalExpression) String() string {
	return fmt.Sprintf("(%s ? %s : %s)", ce.condition, ce.thenExp, ce.elseExp)
}

//赋值表达式
type AssignExpression struct {
	name NameExpression
	value NameExpression
}

func (ae AssignExpression) String() string {
	return fmt.Sprintf("(%s = %s)", ae.name, ae.value)
}

//括号表达式
//type ParenthesesExpression struct {

//}*/

type OperatorExpression struct {
	left NameExpression
	operator token.Token
	right NameExpression
}

func (oe OperatorExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", oe.left, oe.operator, oe.right)
}
