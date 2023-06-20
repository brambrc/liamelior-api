package Model

import "mime/multipart"

type AuthenticationInput struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
    Email    string `form:"email" json:"email" binding:"required"`
    Name     string `form:"name" json:"name" binding:"required"`
    Role     string `form:"role" json:"role" binding:"required"`
	Avatar   *multipart.FileHeader `form:"avatar" json:"avatar" binding:"required"`

}
