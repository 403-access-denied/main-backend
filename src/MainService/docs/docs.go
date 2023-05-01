// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api-call/generatePosterInfo": {
            "get": {
                "description": "Generates info for a poster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApiCall"
                ],
                "summary": "Generate poster info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Image Url",
                        "name": "image_url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.GeneratedPosterInfoView"
                        }
                    }
                }
            }
        },
        "/posters": {
            "get": {
                "description": "Retrieves a list of all posters, sorted and paginated according to the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posters"
                ],
                "summary": "Get a list of all posters",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "asc",
                        "description": "Sort direction",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "created_at",
                        "description": "Sort by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search phrase",
                        "name": "search_phrase",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "both",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Time start",
                        "name": "time_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Time end",
                        "name": "time_end",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Only rewards",
                        "name": "only_rewards",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Longitude",
                        "name": "lon",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "TagIds",
                        "name": "tag_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/View.PosterView"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a poster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posters"
                ],
                "summary": "Create a poster",
                "parameters": [
                    {
                        "description": "Poster",
                        "name": "poster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.CreatePosterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.PosterView"
                        }
                    }
                }
            }
        },
        "/posters/{id}": {
            "get": {
                "description": "Retrieves a poster by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posters"
                ],
                "summary": "Get a poster by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poster ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.PosterView"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a poster by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posters"
                ],
                "summary": "Delete a poster by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poster ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "patch": {
                "description": "Updates a poster by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posters"
                ],
                "summary": "Update a poster by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poster ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Poster",
                        "name": "poster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.UpdatePosterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.PosterView"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "description": "Retrieves Tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Get a Tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/View.CategoryView"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Tag by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Create a Tag by ID",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.CreateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.CategoryView"
                        }
                    }
                }
            }
        },
        "/tags/{id}": {
            "get": {
                "description": "Retrieves a Tag by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Get a Tag by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.CategoryView"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a Tag by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Delete a Tag by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "patch": {
                "description": "Updates a Tag by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Update a Tag by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tag",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.UpdateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.CategoryView"
                        }
                    }
                }
            }
        },
        "/users/auth/google/callback": {
            "get": {
                "description": "google callback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "google callback",
                "responses": {}
            }
        },
        "/users/auth/google/login": {
            "get": {
                "description": "login user with oauth2",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user with oauth2",
                "responses": {}
            }
        },
        "/users/auth/otp/login": {
            "post": {
                "description": "login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "Verify OTP",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.UserView"
                        }
                    }
                }
            }
        },
        "/users/auth/otp/send": {
            "post": {
                "description": "send otp to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "send otp to user",
                "parameters": [
                    {
                        "description": "Send OTP",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UseCase.SendOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/View.MessageView"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "DTO.AddressDTO": {
            "type": "object",
            "required": [
                "city",
                "latitude",
                "longitude",
                "province"
            ],
            "properties": {
                "address_detail": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 5
                },
                "city": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "province": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                }
            }
        },
        "DTO.PosterDTO": {
            "type": "object",
            "required": [
                "alert",
                "status",
                "title",
                "user_id"
            ],
            "properties": {
                "alert": {
                    "type": "boolean"
                },
                "award": {
                    "type": "number"
                },
                "description": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 5
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "lost",
                        "found"
                    ]
                },
                "tel_id": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "user_phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11
                }
            }
        },
        "Model.Address": {
            "type": "object",
            "properties": {
                "address_detail": {
                    "type": "string"
                },
                "address_id": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "province": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "Model.Category": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "posters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Poster"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "Model.Conversation": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Message"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_1_id": {
                    "type": "integer"
                },
                "user_2_id": {
                    "type": "integer"
                }
            }
        },
        "Model.Image": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "image_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "Model.MarkedPoster": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "poster": {
                    "$ref": "#/definitions/Model.Poster"
                },
                "poster_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "Model.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "conversation_id": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "receiver_id": {
                    "type": "integer"
                },
                "sender_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "Model.Poster": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Address"
                    }
                },
                "award": {
                    "type": "number"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Category"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description;": {
                    "type": "string"
                },
                "has_alert": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Image"
                    }
                },
                "status": {
                    "$ref": "#/definitions/Model.PosterStatus"
                },
                "telegram_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/Model.User"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_phone": {
                    "type": "string"
                }
            }
        },
        "Model.PosterStatus": {
            "type": "string",
            "enum": [
                "lost",
                "found"
            ],
            "x-enum-varnames": [
                "Lost",
                "Found"
            ]
        },
        "Model.User": {
            "type": "object",
            "properties": {
                "conversations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Conversation"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "marked_posters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.MarkedPoster"
                    }
                },
                "posters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Poster"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "UseCase.CreatePosterRequest": {
            "type": "object",
            "required": [
                "categories",
                "img_urls"
            ],
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/DTO.AddressDTO"
                    }
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "img_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "poster": {
                    "$ref": "#/definitions/DTO.PosterDTO"
                }
            }
        },
        "UseCase.CreateTagRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 1
                }
            }
        },
        "UseCase.SendOTPRequest": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 11
                }
            }
        },
        "UseCase.UpdatePosterRequest": {
            "type": "object",
            "required": [
                "categories",
                "img_urls"
            ],
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/DTO.AddressDTO"
                    }
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "img_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "poster": {
                    "$ref": "#/definitions/DTO.PosterDTO"
                }
            }
        },
        "UseCase.UpdateTagRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 1
                }
            }
        },
        "UseCase.VerifyOTPRequest": {
            "type": "object",
            "required": [
                "otp",
                "username"
            ],
            "properties": {
                "otp": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 11
                }
            }
        },
        "View.CategoryView": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "View.GeneratedPosterInfoView": {
            "type": "object",
            "properties": {
                "colors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "titles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "View.MessageView": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "View.PosterView": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Address"
                    }
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Category"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Model.Image"
                    }
                },
                "status": {
                    "$ref": "#/definitions/Model.PosterStatus"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "type": "integer"
                }
            }
        },
        "View.UserView": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
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
	Title:            "Swagger Documentation for Golang web API(Gin framework)",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
