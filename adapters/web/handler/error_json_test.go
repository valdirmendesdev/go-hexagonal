package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}