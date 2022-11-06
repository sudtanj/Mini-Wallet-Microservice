package wallet

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateSecureToken(t *testing.T) {
	generatedToken := GenerateSecureToken(12)
	fmt.Println("Generate token with pattern [a-z0-9]")
	assert.MatchRegex(t, generatedToken, "^[a-z0-9]*$")
	fmt.Println("Check if token size is 24")
	assert.Equal(t, len(generatedToken), 24)
	fmt.Println("Check if token not have any capital letter")
	assert.NotMatchRegex(t, generatedToken, "^[A-Z]*$")
}
