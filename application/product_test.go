package application_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.Error(t, err, "the price must be greater than zero to enable the product")
	require.Equal(t, application.ENABLED, product.Status)
}
