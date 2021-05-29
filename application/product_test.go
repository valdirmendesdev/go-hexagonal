package application_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product",err)
	require.Equal(t, application.ENABLED, product.Status)
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to disable the product",err)
	require.Equal(t, application.DISABLED, product.Status)
}
