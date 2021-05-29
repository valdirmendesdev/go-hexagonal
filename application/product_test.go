package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/go-hexagonal/application"
	"testing"
)

func createNewProduct(id, name, status string, price float64 ) application.Product {
	return application.Product{
		ID:     id,
		Name:   name,
		Price:  price,
		Status: status,
	}
}

func TestProduct_Enable(t *testing.T) {
	product := createNewProduct("", "", application.DISABLED, 10)

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.EqualError(t, err,"the price must be greater than zero to enable the product")
	require.Equal(t, application.ENABLED, product.Status)
}

func TestProduct_Disable(t *testing.T) {
	product := createNewProduct("","", application.ENABLED, 0)

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10
	err = product.Disable()
	require.EqualError(t, err, "the price must be zero in order to disable the product")
	require.Equal(t, application.DISABLED, product.Status)
}

func TestProduct_IsValid(t *testing.T) {
	product := createNewProduct(uuid.NewV4().String(), "hello", application.DISABLED, 10)

	isValid, err := product.IsValid()
	require.Nil(t, err)
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
}
