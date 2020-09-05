package gotest

import (
	"testing"
)

func BenchmarkGenerateKey(b *testing.B) {
	for i := 0; i < b.N; i++ { // use bN for looping
		_, err := generateKey()
		if err != nil {
			b.Error("generate key failed")
		}
	}
}
