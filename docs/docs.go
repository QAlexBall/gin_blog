// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api-user/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get User List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/api-user/user/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get User By ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Update User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
