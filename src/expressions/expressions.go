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

