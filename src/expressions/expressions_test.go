package expressions

import (
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"testing"
	"context"
	"fmt"
	"github.com/ckLXHL/gotoy/src/token"
)

func Test(t *testing.T) {
	t.Run("NameExpression", test.Case(func(ctx context.Context) {
		src := "12345"
		dst := NameExpression{src}
		must.Equal(src, fmt.Sprint(dst))
	}))
/*
	t.Run("ConditionalExpression", test.Case(func(ctx context.Context) {
		src1, src2, src3 := "12345", "7899", "11111"
		dst := ConditionalExpression{
			NameExpression{src1},
			NameExpression{src2},
			NameExpression{src3},
		}
		must.Equal(fmt.Sprintf("(%s ? %s : %s)", src1, src2, src3), fmt.Sprint(dst))
	}))
	t.Run("AssignExpression", test.Case(func(ctx context.Context) {
		src1, src2 := "1234", "5678"
		dst := AssignExpression{
			NameExpression{src1},
			NameExpression{src2},
		}
		must.Equal(fmt.Sprintf("(%s = %s)", src1, src2), fmt.Sprint(dst))
	}))*/
	t.Run("OperatorExpression", test.Case(func(ctx context.Context) {
		src1, src2, src3 := "123", "+", "456"
		dst := OperatorExpression{
			NameExpression{src1},
			token.Token{token.PLUS},
			NameExpression{src3},
		}
		must.Equal(fmt.Sprintf("(%s %s %s)", src1, src2, src3), fmt.Sprint(dst))
	}))
}