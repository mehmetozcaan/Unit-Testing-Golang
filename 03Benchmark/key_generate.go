package gotest

import (
	"crypto/rand"
	"fmt"
)

func generateKey() ([]byte, error) {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("Error detail: %v", err)
	}
	return key, nil
}
