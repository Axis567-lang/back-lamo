package token

import (
	"testing"
)

func TestMakeToken(t *testing.T) {
	const tokens = 10
	var tk []Token
	for i := 0; i < tokens; i++ {
		println(i)
		tk = append(tk, MakeToken(Size))
		println(tk[i].String())
	}

	for i := 0; i < tokens-1; i++ {
		for j := i + 1; j < tokens; j++ {
			println(i, j)
			if tk[i].String() == tk[j].String() {
				t.Fatalf("duplicated token")
			}
		}
	}

}
