package types

import (
	"fmt"
	"testing"
)

func TestNewPoolID(t *testing.T) {
	pid := NewPoolID()
	fmt.Printf("Pool ID: %v\n", pid)
}

func TestNewUserID(t *testing.T) {
	pid := NewUserID()
	fmt.Printf("Pool ID: %v\n", pid)
}