package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min := int64(0)
	max := int64(1000)
	randomInt := RandomInt(min, max)

	require.NotEmpty(t, randomInt)
	require.GreaterOrEqual(t, randomInt, min)
	require.LessOrEqual(t, randomInt, max)
}

func TestRandomString(t *testing.T) {
	randomStringLength := RandomInt(1, 10)
	randomString := RandomString(int(randomStringLength))

	require.Len(t, randomString, int(randomStringLength))
	require.NotEmpty(t, randomString)
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()

	require.NotEmpty(t, owner)
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()

	require.NotEmpty(t, money)
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	validCurrencies := []string{"EUR", "USD", "CAD"}

	require.NotEmpty(t, currency)
	require.Contains(t, validCurrencies, currency)
}
