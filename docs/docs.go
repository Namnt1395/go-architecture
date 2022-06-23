// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nguyen Thanh Nam",
            "email": "namnguyenthanh024@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/capacities": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Capacity"
                ],
                "summary": "Get list capacity",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "created_at desc",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "is_active",
                        "name": "is_active",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "capacity_group_ids[]",
                        "name": "capacity_group_ids[]",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListCapacityOutput"
                        }
                    },
                    "500": {
                        "description": "{\"error_code\": \"\u003cMã lỗi\u003e\", \"error_msg\": \"\u003cNội dung lỗi\u003e\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new capacity",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Capacity"
                ],
                "summary": "Create new capacity",
                "parameters": [
                    {
                        "description": "JSON body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CapacityInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CapacityOutput"
                        }
                    },
                    "500": {
                        "description": "{\"error_code\": \"\u003cMã lỗi\u003e\", \"error_msg\": \"\u003cNội dung lỗi\u003e\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/v1/capacities/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Capacity"
                ],
                "summary": "Get info capacity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CapacityOutput"
                        }
                    },
                    "500": {
                        "description": "{\"error_code\": \"\u003cMã lỗi\u003e\", \"error_msg\": \"\u003cNội dung lỗi\u003e\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update capacity",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Capacity"
                ],
                "summary": "Update capacity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CapacityInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CapacityOutput"
                        }
                    },
                    "500": {
                        "description": "{\"error_code\": \"\u003cMã lỗi\u003e\", \"error_msg\": \"\u003cNội dung lỗi\u003e\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Capacity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Capacity"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "{\"error_code\": \"\u003cMã lỗi\u003e\", \"error_msg\": \"\u003cNội dung lỗi\u003e\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CapacityData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "dto.CapacityInput": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CapacityOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.CapacityData"
                }
            }
        },
        "dto.ListCapacityOutput": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CapacityData"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8099",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Sample Service API By Layer",
	Description: "This is Sample Service API By Layer.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
