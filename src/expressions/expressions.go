package expressions

import "fmt"

//expression格式化接口
type ExpressionIf fmt.Stringer

//名称表达式
type NameExpression struct {
	name string
}

func (ne NameExpression) String() string {
	return ne.name
}

//条件表达式
type ConditionalExpression struct{
	condition, thenArm, elseArm ExpressionIf
}

func (ce ConditionalExpression) String() string {
	return fmt.Sprintf("(%s ? %s : %s)", ce.condition, ce.thenArm, ce.elseArm)
}

//赋值表达式
type AssignExpression struct {
	name string
	right NameExpression
}

func (ae AssignExpression) String() string {
	return fmt.Sprintf("(%s = %s)", ae.name, ae.right)
}