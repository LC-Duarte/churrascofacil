// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/calcular": {
            "get": {
                "description": "Calcula a quantidade de carne necessária para um churrasco",
                "produces": [
                    "application/json"
                ],
                "summary": "Calcular quantidade de carne",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Quantidade de pessoas",
                        "name": "pessoas",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "number"
                            }
                        }
                    }
                }
            }
        },
        "/carnes": {
            "get": {
                "description": "Retorna uma lista de tipos de carne disponíveis",
                "produces": [
                    "application/json"
                ],
                "summary": "Listar tipos de carne",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/cortes": {
            "get": {
                "description": "Retorna uma lista de cortes de churrasco para um determinado tipo de carne",
                "produces": [
                    "application/json"
                ],
                "summary": "Listar cortes de churrasco",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tipo de carne (Frango, Boi, Porco, Carneiro)",
                        "name": "tipo",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Churrasco Fácil API",
	Description:      "API para listar e calcular quantidades de carne para um churrasco fácil.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
