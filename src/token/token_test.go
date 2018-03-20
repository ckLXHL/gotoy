package token

import (
"github.com/modern-go/test"
"github.com/modern-go/test/must"
"testing"
"context"
	"fmt"
)

func Test(t *testing.T) {
	t.Run("tokenType", test.Case(func(ctx context.Context) {
		src := []string{"+", "-", "*", "/", "~", "!", "(", ")", "="}
		srcNames := []TokenType{PLUS, MINUS, ASTERISK, SLASH, TILDE, BANG, LEST_PAREN, RIGHT_PAREN, ASSIGNMENG}

		for i, v := range src {
			fetch, err := Punctuator[srcNames[i]]
			must.Equal(true, err)
			must.Equal(v, fetch)
		}
	}))
	t.Run("Token", test.Case(func(ctx context.Context) {
		dst := []string{"+", "-", "*", "/", "~", "!", "(", ")", "="}
		src := []TokenType{PLUS, MINUS, ASTERISK, SLASH, TILDE, BANG, LEST_PAREN, RIGHT_PAREN, ASSIGNMENG}

		for i, v := range src {
			fetch := Token{v}
			must.Equal(dst[i], fmt.Sprint(fetch))
		}
	}))

}