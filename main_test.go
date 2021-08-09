package main_test

import (
	"codacy-test/main"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringPtr(t *testing.T) {
	stringToCheck := "x"
	res := main.StringPtr(stringToCheck)
	require.Equal(t, &stringToCheck, res)
}
