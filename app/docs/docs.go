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
        "/app/sms_country_list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "获取短信支持的国家 code 列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "国家列表",
                        "schema": {
                            "$ref": "#/definitions/response.SmsCountryListResp"
                        }
                    }
                }
            }
        },
        "/app/test": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "app test 接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/room/create_room": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "创建房间",
                "parameters": [
                    {
                        "description": "创建房间",
                        "name": "CreateRoomReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room.CreateRoomReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/room.CreateRoomResp"
                        }
                    }
                }
            }
        },
        "/room/get_room_base_info": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "获取房间基础信息",
                "parameters": [
                    {
                        "description": "房间基础信息",
                        "name": "RoomBasenfoReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room.RoomBasenfoReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/room.RoomBasenfoResp"
                        }
                    }
                }
            }
        },
        "/room/room_list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "房间列表",
                "parameters": [
                    {
                        "description": "房间列表",
                        "name": "RoomListReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room.RoomListReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/room.RoomListResp"
                        }
                    }
                }
            }
        },
        "/sys/get_sys_config": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sys"
                ],
                "summary": "返回前端使用的系统配置",
                "parameters": [
                    {
                        "description": "获取系统配置",
                        "name": "GetSysConfigReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sys.GetSysConfigReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/sys.GetSysConfigResp"
                        }
                    }
                }
            }
        },
        "/user/account_verify_code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "验证登录账号的验证码是否有效",
                "parameters": [
                    {
                        "description": "验证登录账号的验证码",
                        "name": "ChangePassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AccountVerifyCodeReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/user.AccountVerifyCodeResp"
                        }
                    }
                }
            }
        },
        "/user/change_password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用户密码",
                "parameters": [
                    {
                        "description": "修改密码",
                        "name": "ChangePassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ChangePasswordReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/user.ChangePasswordResp"
                        }
                    }
                }
            }
        },
        "/user/get_rc_user_token": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取融云用户 token",
                "parameters": [
                    {
                        "description": "获取user token",
                        "name": "GetUserToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rc.GetUserTokenReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/rc.GetUserTokenResp"
                        }
                    }
                }
            }
        },
        "/user/get_user_info": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "description": "获取用户休息",
                        "name": "ResetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.GetUserInfoReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/user.GetUserInfoResp"
                        }
                    }
                }
            }
        },
        "/user/get_verify_code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取验证码",
                "parameters": [
                    {
                        "description": "获取验证码",
                        "name": "GetVerifyCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.GetVerifyCodeReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/user.GetVerifyCodeResp"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "手机\\邮箱验证码登录",
                "parameters": [
                    {
                        "description": "用户登录",
                        "name": "LoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserLoginReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/user.UserLoginResp"
                        }
                    }
                }
            }
        },
        "/user/reset_password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "重置用户密码",
                "parameters": [
                    {
                        "description": "重置密码",
                        "name": "ResetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ResetPasswordReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "登录Token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/user.ResetPasswordResp"
                        }
                    }
                }
            }
        },
        "/user/third_login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "第三方登录登录",
                "parameters": [
                    {
                        "description": "第三方用户登录",
                        "name": "ThirdLoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ThirdUserLoginReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "请求ID",
                        "name": "request_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/user.ThirdUserLoginResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "_struct.MobilePhoneAccountInfo": {
            "type": "object",
            "properties": {
                "area_code": {
                    "description": "区号",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                }
            }
        },
        "rc.GetUserTokenReq": {
            "type": "object"
        },
        "rc.GetUserTokenResp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_sys_id": {
                    "description": "uid 用户id",
                    "type": "string"
                }
            }
        },
        "response.SmsCountryListResp": {
            "type": "object",
            "properties": {
                "area_code": {
                    "description": "国家手机区号",
                    "type": "string"
                },
                "code": {
                    "description": "国家 code",
                    "type": "string"
                },
                "en_name": {
                    "description": "英文名称",
                    "type": "string"
                },
                "flag_url": {
                    "description": "国旗 url",
                    "type": "string"
                },
                "name": {
                    "description": "国家名称",
                    "type": "string"
                }
            }
        },
        "room.CreateRoomReq": {
            "type": "object"
        },
        "room.CreateRoomResp": {
            "type": "object",
            "properties": {
                "info": {
                    "description": "房间的基本信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/room.RoomBaseInfo"
                        }
                    ]
                }
            }
        },
        "room.RoomBaseInfo": {
            "type": "object",
            "properties": {
                "cover": {
                    "description": "房间封面",
                    "type": "string"
                },
                "im_group_id": {
                    "description": "IM群组ID",
                    "type": "integer"
                },
                "live_status": {
                    "description": "直播状态 open:直播中 close:未直播",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "notice": {
                    "description": "房间公告",
                    "type": "string"
                },
                "status": {
                    "description": "房间状态 open:正常 close:禁用",
                    "type": "string"
                },
                "sys_id": {
                    "description": "房间 SysID",
                    "type": "string"
                }
            }
        },
        "room.RoomBasenfoReq": {
            "type": "object",
            "required": [
                "room_sys_id"
            ],
            "properties": {
                "room_sys_id": {
                    "description": "房间的 sysID",
                    "type": "string"
                }
            }
        },
        "room.RoomBasenfoResp": {
            "type": "object",
            "properties": {
                "info": {
                    "description": "房间的基本信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/room.RoomBaseInfo"
                        }
                    ]
                }
            }
        },
        "room.RoomListReq": {
            "type": "object",
            "required": [
                "page",
                "page_size"
            ],
            "properties": {
                "page": {
                    "description": "第几页，从 1 开始",
                    "type": "integer"
                },
                "page_size": {
                    "description": "每页的大小",
                    "type": "integer"
                }
            }
        },
        "room.RoomListResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/room.RoomBaseInfo"
                    }
                },
                "total": {
                    "description": "房间总条数",
                    "type": "integer"
                }
            }
        },
        "sys.GetSysConfigReq": {
            "type": "object"
        },
        "sys.GetSysConfigResp": {
            "type": "object",
            "properties": {
                "agora_app_id": {
                    "description": "声望 app id",
                    "type": "string"
                },
                "rc_app_key": {
                    "description": "融云 app key",
                    "type": "string"
                }
            }
        },
        "user.AccountBindInfo": {
            "type": "object",
            "properties": {
                "mobile_phone": {
                    "$ref": "#/definitions/_struct.MobilePhoneAccountInfo"
                }
            }
        },
        "user.AccountVerifyCodeReq": {
            "type": "object",
            "required": [
                "code",
                "identity_type"
            ],
            "properties": {
                "code": {
                    "description": "验证码",
                    "type": "string"
                },
                "identity_type": {
                    "description": "登录类型，mobile_phone、email 目前只有 mobile_phone",
                    "type": "string"
                }
            }
        },
        "user.AccountVerifyCodeResp": {
            "type": "object",
            "properties": {
                "is_verify": {
                    "description": "是否验证通过",
                    "type": "boolean"
                },
                "verify_hash": {
                    "description": "本次验证的 hash, 后续重置密码使用",
                    "type": "string"
                }
            }
        },
        "user.ChangePasswordReq": {
            "type": "object",
            "required": [
                "identity_type",
                "new_password",
                "password"
            ],
            "properties": {
                "identity_type": {
                    "description": "登录类型，mobile_phone、email 目前只有 mobile_phone",
                    "type": "string"
                },
                "new_password": {
                    "description": "新密码 前端 必须进行 sha256 hash 后传过来 【小写】",
                    "type": "string"
                },
                "password": {
                    "description": "登录密码，前端 必须进行 sha256 hash 后传过来 【小写】",
                    "type": "string"
                }
            }
        },
        "user.ChangePasswordResp": {
            "type": "object"
        },
        "user.GetUserInfoReq": {
            "type": "object",
            "properties": {
                "return_account_bind_list": {
                    "description": "是否返回账号绑定列表，默认不返回",
                    "type": "boolean"
                },
                "return_fans_data": {
                    "description": "是否返回粉丝数据",
                    "type": "boolean"
                }
            }
        },
        "user.GetUserInfoResp": {
            "type": "object",
            "properties": {
                "account_bind_info": {
                    "description": "绑定账号信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.AccountBindInfo"
                        }
                    ]
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "birthday": {
                    "description": "生日年-月-日",
                    "type": "string"
                },
                "fan_data": {
                    "description": "粉丝数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.UserFanData"
                        }
                    ]
                },
                "introduction": {
                    "description": "个人简介",
                    "type": "string"
                },
                "nick_name": {
                    "description": "昵称",
                    "type": "string"
                },
                "room_sys_id": {
                    "description": "Room Id",
                    "type": "string"
                },
                "sex": {
                    "description": "性别 male:男;female:女; unknown:未知 保密",
                    "type": "string"
                },
                "user_sys_id": {
                    "description": "uid 用户id",
                    "type": "string"
                }
            }
        },
        "user.GetVerifyCodeReq": {
            "type": "object",
            "required": [
                "account",
                "area_code",
                "identity_type",
                "scene"
            ],
            "properties": {
                "account": {
                    "description": "登录账号",
                    "type": "string"
                },
                "area_code": {
                    "description": "区号,如果是手机号 传 区号，如果是邮箱 传 邮箱服务商 gmail、qq、outlook 等",
                    "type": "string"
                },
                "identity_type": {
                    "description": "登录类型，mobile_phone、email 目前只有 mobile_phone",
                    "type": "string"
                },
                "scene": {
                    "description": "场景 login 登录、reset_password 重置密码",
                    "type": "string"
                }
            }
        },
        "user.GetVerifyCodeResp": {
            "type": "object",
            "properties": {
                "is_new": {
                    "description": "验证码\nCode string ` + "`" + `json:\"code\"` + "`" + `\n是否是新用户",
                    "type": "boolean"
                },
                "ttl": {
                    "description": "验证码有效时长",
                    "type": "integer"
                }
            }
        },
        "user.ResetPasswordReq": {
            "type": "object",
            "required": [
                "identity_type",
                "new_password",
                "verify_hash"
            ],
            "properties": {
                "identity_type": {
                    "description": "登录类型，mobile_phone、email 目前只有 mobile_phone",
                    "type": "string"
                },
                "new_password": {
                    "description": "新密码 前端 必须进行 sha256 hash 后传过来 【小写】",
                    "type": "string"
                },
                "verify_hash": {
                    "description": "验证码 hash, 会在验证码验证没问题后返回，代表验证成功",
                    "type": "string"
                }
            }
        },
        "user.ResetPasswordResp": {
            "type": "object"
        },
        "user.ThirdUserLoginReq": {
            "type": "object",
            "required": [
                "identity_type",
                "third_token"
            ],
            "properties": {
                "firebase_token": {
                    "description": "firebase token",
                    "type": "string"
                },
                "identity_type": {
                    "description": "@Summary 第三方登录类型 apple、google",
                    "type": "string"
                },
                "third_token": {
                    "description": "第三方登录 token",
                    "type": "string"
                }
            }
        },
        "user.ThirdUserLoginResp": {
            "type": "object",
            "properties": {
                "is_new": {
                    "description": "是否是新用户",
                    "type": "boolean"
                },
                "token": {
                    "description": "用户登录成功的 token",
                    "type": "string"
                }
            }
        },
        "user.UserFanData": {
            "type": "object",
            "properties": {
                "each_other_count": {
                    "description": "相互关注数量",
                    "type": "integer"
                },
                "fan_count": {
                    "description": "粉丝数",
                    "type": "integer"
                },
                "follow_count": {
                    "description": "关注数",
                    "type": "integer"
                }
            }
        },
        "user.UserLoginReq": {
            "type": "object",
            "required": [
                "account",
                "area_code",
                "identity_type",
                "login_type"
            ],
            "properties": {
                "account": {
                    "description": "登录账号",
                    "type": "string"
                },
                "area_code": {
                    "description": "区号,如果是手机号 传 区号，如果是邮箱 传 邮箱服务商 gmail、qq、outlook 等",
                    "type": "string"
                },
                "code": {
                    "description": "验证码",
                    "type": "string"
                },
                "identity_type": {
                    "description": "登录类型，mobile_phone、email 目前只有 mobile_phone",
                    "type": "string"
                },
                "login_type": {
                    "description": "登录方式 password 密码登录、verify_code 验证码登录",
                    "type": "string"
                },
                "password": {
                    "description": "登录密码，前端 必须进行 sha256 hash 后传过来 【小写】\n1. 如果账号存在的时候登录，密码 或者验证码必须有一个\n2. 如果获取 验证码登录的时候，获取验证码返回的是新用户，需要设置密码，也是使用该字段",
                    "type": "string"
                }
            }
        },
        "user.UserLoginResp": {
            "type": "object",
            "properties": {
                "is_new": {
                    "description": "是否是新用户",
                    "type": "boolean"
                },
                "token": {
                    "description": "用户登录成功的 token",
                    "type": "string"
                },
                "user_sys_id": {
                    "description": "uid 用户id",
                    "type": "string"
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}