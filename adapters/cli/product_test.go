package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/go-hexagonal/adapters/cli"
	"github.com/valdirmendesdev/go-hexagonal/application"
	mock_application "github.com/valdirmendesdev/go-hexagonal/application/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productID := "abc"
	productName := "Test Product"
	productPrice := 25.99
	productStatus := application.ENABLED

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	testsTable := []struct {
		action string
		result string
	}{
		{
			action: "create",
			result: fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s",
				productID,
				productName,
				productPrice,
				productStatus),
		},
		{
			action: "enable",
			result: fmt.Sprintf("Product %s has been enabled.", productName),
		},
		{
			action: "disable",
			result: fmt.Sprintf("Product %s has been disabled.", productName),
		},
		{
			action: "get",
			result: fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
				productID,
				productName,
				productPrice,
				productStatus),
		},
	}

	for _, test := range testsTable {
		result, err := cli.Run(service, test.action, productID, productName, productPrice)
		require.Nil(t, err)
		require.Equal(t, test.result, result)
	}
}
