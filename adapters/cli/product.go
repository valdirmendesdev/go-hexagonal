package cli

import (
	"fmt"
	"github.com/valdirmendesdev/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action, productID, productName string, price float64) (string, error) {

	var result string

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		res, err := service.Enable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		res, err := service.Disable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	}

	return result, nil
}
