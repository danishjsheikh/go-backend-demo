package core

import (
	"fmt"
	"strings"

	"github.com/danishjsheikh/backend-go/app/common"
	"github.com/danishjsheikh/backend-go/app/models"
)

func GetOrderByCode(correlationId string, orderCode string) string {
	fmt.Printf("Your correlationId is %v", correlationId)
	return fmt.Sprintf("This is your order Code: %v", orderCode)
}

func CreateOrder(request models.CreateOrderRequest) (string, error) {

	if errs := common.CustomValidatorGlobal.Validate(common.CustomValidatorGlobal.Validator, request); len(errs) > 0 && errs[0].HasError {
		errorMessages := make([]string, 0)

		for _, err2 := range errs {
			errorMessages = append(errorMessages, fmt.Sprintf("%s field has failed. Validation is: %s", err2.Field, err2.Tag))
		}

		return "", fmt.Errorf("%v", strings.Join(errorMessages, "\n"))
	}

	return "Order created successfully!", nil
}
