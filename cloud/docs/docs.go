// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
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
        "/exercises": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exercises"
                ],
                "summary": "Create a couple of exercises",
                "parameters": [
                    {
                        "description": "Exercises",
                        "name": "exercises",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/exercise.ExerciseRecord"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "exercise.ExerciseData": {
            "type": "object",
            "required": [
                "finished_at",
                "started_at",
                "weight"
            ],
            "properties": {
                "finished_at": {
                    "type": "integer"
                },
                "started_at": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "exercise.ExerciseRecord": {
            "type": "object",
            "required": [
                "data",
                "origin_id",
                "user_id"
            ],
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/exercise.ExerciseData"
                    }
                },
                "origin_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Smart-Gym API",
	Description:      "This is the strongest server ever.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
