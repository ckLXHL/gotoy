package operation

import (
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"testing"
	"context"
)

func Test(t *testing.T) {
	t.Run("1＋1", test.Case(func(ctx context.Context) {
		src := 1+1
		dst := Add(1, 1)
		must.Equal(src, dst)
	}))
}