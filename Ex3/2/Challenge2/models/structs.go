package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,checkstring"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,checkstring"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role" example:"user"`
	Email       string `json:"email" example:"email"`
	TokenString string `json:"token" example:"Token jwt"`
}

type UserSignUp struct {
	Name     string `json:"name" example:"Tên"`
	Email    string `json:"email" example:"Email"`
	Password string `json:"password" example:"Mật khẩu"`
}

type ExampleSuccess struct {
	Message string `json:"message" example:"Thông báo đã thực hiện thành công..."`
}

type ExampleFailure struct {
	Error string `json:"error" example:"Lỗi gì đó ..."`
}

type UserSignIn struct {
	Email    string `json:"email" example:"Email"`
	Password string `json:"password" example:"Mật khẩu"`
}

type UserAdd struct {
	Name     string `json:"name" example:"Tên"`
	Email    string `json:"email" example:"Email"`
	Password string `json:"password" example:"Mật khẩu"`
	Role 	 string `json:"role" example:"admin/user"`
}

