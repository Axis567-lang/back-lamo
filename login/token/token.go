package token

import "math/rand"

func MakeToken(n int) Token {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return Token{Value: string(b)}
}

type Token struct {
	Value string
}

func (t Token) String() string {
	return t.Value
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!#$%&=^*.-_|"
const Size = 32
