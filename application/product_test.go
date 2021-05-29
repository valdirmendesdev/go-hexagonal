package application_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"testing"
)

func createNewProduct(name string, price float64 ) *application.Product {
	product := application.NewProduct()
	product.Name = name
	product.Price = price
	return product
}

func TestProduct_Enable(t *testing.T) {
	product := createNewProduct("", 10)

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.EqualError(t, err,"the price must be greater than zero to enable the product")
	require.Equal(t, application.ENABLED, product.Status)
}

func TestProduct_Disable(t *testing.T) {
	product := createNewProduct("", 0)
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10
	err = product.Disable()
	require.EqualError(t, err, "the price must be zero in order to disable the product")
	require.Equal(t, application.DISABLED, product.Status)
}

func TestProduct_IsValid(t *testing.T) {
	product := createNewProduct("hello", 10)

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
	require.True(t, isValid)

	product.Status = "invalid"
	isValid, err = product.IsValid()
	require.EqualError(t, err, "the status must be enabled or disabled")
	require.False(t, isValid)

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1
	isValid, err = product.IsValid()
	require.EqualError(t, err, "the price must be greater or equal to zero")
	require.False(t, isValid)

	product = &application.Product{}
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "ID: Missing required field;Name: non zero value required")
}
