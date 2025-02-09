package main

import (
	"log"

	"github.com/danishjsheikh/go-backend-demo/app/common"
	"github.com/danishjsheikh/go-backend-demo/app/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @title Demo Backend
// @version 1
// @Description Demo Backend For MCP

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	app := gin.Default()

	var validate = validator.New()
	common.CustomValidatorGlobal = &common.CustomValidator{
		Validator: validate,
	}

	err := common.AddOldAgeCustomValidation(common.CustomValidatorGlobal)
	if err != nil {
		return
	}

	routes.Init(app)

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
