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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "checks the user and returns tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "User Information to log in",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ReqLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns access and refresh tokens",
                        "schema": {
                            "$ref": "#/definitions/auth.Tokens"
                        }
                    },
                    "400": {
                        "description": "You did something wrong",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "if Access token fails it will returns this",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "removes refresh token gets token from header",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "log outs user",
                "operationId": "logout",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "some thing wrong with what you sent",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Invalid token in header",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/auth/refreshtoken": {
            "post": {
                "description": "gives new access token through refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refresh token",
                "operationId": "refresh",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "some thing wrong with what you sent",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Invalid token in header",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registers user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register User",
                "operationId": "register",
                "parameters": [
                    {
                        "description": "User information to create it",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ReqCreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Invalid inputs can result to ",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/auth/resetpassword": {
            "post": {
                "description": "send info about reserttting poassword to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "resets password",
                "operationId": "reset",
                "parameters": [
                    {
                        "description": "email of the user",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ReqResetPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "some thing wrong with what you sent",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Invalid token in header",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/auth/updatepassword/{email}": {
            "post": {
                "description": "updates password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "update password",
                "operationId": "updatepassword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email of the user",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "email of the user",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ReqUpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "some thing wrong with what you sent",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Invalid token in header",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Something went wrong in server",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.ReqCreateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.ReqLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.ReqResetPassword": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.ReqUpdatePassword": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9999",
	BasePath:         "/localeats.uz",
	Schemes:          []string{},
	Title:            "LocalEats API",
	Description:      "LocalEats is a program to order local and homemade food with quality and precise delivery.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
