// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "post": {
                "description": "Creates a new order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Correlation ID",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Order Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            }
        },
        "/orders/{orderCode}": {
            "get": {
                "description": "Retrieves an order by order code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get Order by Code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Correlation ID",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order Code",
                        "name": "orderCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Update Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Correlation ID",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order Code",
                        "name": "orderCode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Order",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an order",
                "tags": [
                    "Orders"
                ],
                "summary": "Delete Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Code",
                        "name": "orderCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "return list of all users from the database",
                "tags": [
                    "Users"
                ],
                "summary": "return list of all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "create new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "put": {
                "description": "Updates an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated User",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user",
                "tags": [
                    "Users"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateOrderRequest": {
            "description": "Request about creating Order",
            "type": "object",
            "required": [
                "age",
                "countryCode",
                "shipmentNumber"
            ],
            "properties": {
                "age": {
                    "description": "age to make sure you are young",
                    "type": "integer"
                },
                "countryCode": {
                    "description": "country code like: tr, us",
                    "type": "string"
                },
                "shipmentNumber": {
                    "description": "shipment no of Order",
                    "type": "string"
                }
            }
        },
        "models.ErrResponse": {
            "description": "User Error Response",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.Order": {
            "description": "Order Schema",
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "countryCode": {
                    "type": "string"
                },
                "shipmentNumber": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "description": "User Schema",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "description": "User Valid Response",
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Demo Backend",
	Description:      "Demo Backend For MCP",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
