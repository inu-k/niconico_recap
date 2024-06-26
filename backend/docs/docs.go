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
        "/history/{date}": {
            "get": {
                "description": "dateが指定された場合はその日の5時から29時までのデータを返す\nstartDateとendDateが指定された場合はstartDateの5時からからendDateの29時までのデータを返す\nstartDateが指定されない場合は1900-01-01, endDateが指定されない場合は現在として扱う",
                "produces": [
                    "application/json"
                ],
                "summary": "指定された日付の視聴履歴を返す",
                "parameters": [
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "endDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.DetailHistory"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/summary": {
            "get": {
                "description": "指定された期間内に視聴された動画のタグのサマリーを返す\nタグは視聴された回数の降順で返される\ndateが指定された場合はその日の5時から29時までのデータを返す\nstartDateとendDateが指定された場合はstartDateの5時からからendDateの29時までのデータを返す\nstartDateが指定されない場合は1900-01-01, endDateが指定されない場合は現在として扱う",
                "produces": [
                    "application/json"
                ],
                "summary": "指定された期間内に視聴された動画のタグのサマリーを返す",
                "parameters": [
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "yyyy-mm-dd",
                        "name": "endDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.TagNameCount"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/videos": {
            "get": {
                "description": "tagが指定された場合はそのタグが含まれる動画を返す\ntitleが指定された場合はそのタイトルが含まれる動画を返す\ntagとtitleが指定された場合はtagの条件が優先される",
                "produces": [
                    "application/json"
                ],
                "summary": "タグやタイトルで動画を検索して返す",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tag",
                        "name": "tag",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.VideoInfoWithoutTags"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/videos/{videoId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "指定されたvideoIdの動画情報を返す",
                "parameters": [
                    {
                        "type": "string",
                        "description": "videoId",
                        "name": "videoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.VideoInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presentation.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.DetailHistory": {
            "type": "object",
            "properties": {
                "thumbnail_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "video_id": {
                    "type": "string"
                },
                "watch_date": {
                    "type": "string"
                }
            }
        },
        "data.TagNameCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "tag_name": {
                    "type": "string"
                }
            }
        },
        "data.VideoInfo": {
            "type": "object",
            "properties": {
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "video_id": {
                    "type": "string"
                }
            }
        },
        "data.VideoInfoWithoutTags": {
            "type": "object",
            "properties": {
                "thumbnail_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "video_id": {
                    "type": "string"
                }
            }
        },
        "presentation.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "niconico_recap_backend API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
