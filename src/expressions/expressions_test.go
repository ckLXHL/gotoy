package expressions

import (
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"testing"
	"context"
	"fmt"
)

func Test(t *testing.T) {
	t.Run("NameExpression", test.Case(func(ctx context.Context) {
		src := "12345"
		dst := NameExpression{src}
		must.Equal(src, fmt.Sprint(dst))
	}))

}