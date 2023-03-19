// Code generated by swaggo/swag. DO NOT EDIT
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
        "/users": {
            "post": {
                "description": "Cria um usuário com base nas informações fornecidas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Cria um usuário",
                "parameters": [
                    {
                        "description": "Informações do usuário",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UserInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.UserInputDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Meu API REST",
	Description:      "Esta é uma API REST de exemplo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}