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
        "/funding": {
            "get": {
                "description": "get all funding",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "funding"
                ],
                "summary": "Get funding list",
                "operationId": "get-funding",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FinalData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                }
            }
        },
        "model.FinalData": {
            "type": "object",
            "properties": {
                "negative": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.FundingCoin"
                    }
                },
                "positive": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.FundingCoin"
                    }
                }
            }
        },
        "types.FundingCoin": {
            "type": "object",
            "properties": {
                "elems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.SuitableCoin"
                    }
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "types.SuitableCoin": {
            "type": "object",
            "properties": {
                "askPrice": {
                    "type": "string"
                },
                "bidPrice": {
                    "type": "string"
                },
                "exchange": {
                    "type": "string"
                },
                "fundingRate": {
                    "type": "number"
                },
                "nextFundingTime": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Funding API",
	Description:      "This is a sample funding server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
