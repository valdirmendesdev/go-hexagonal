package cli

import (
	"fmt"
	"github.com/valdirmendesdev/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action, productID, productName string, price float64) (string, error) {

	var result string

	product, err := service.Create(productName, price)
	if err != nil {
		return "", err
	}

	result = fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s",
				product.GetID(),
				product.GetName(),
				product.GetPrice(),
				product.GetStatus())

	return result, nil
}
