package controller

import (
	"github.com/gin-gonic/gin"
	"go-zero-demo/common"
	models "go-zero-demo/model"
	"time"
)

type UserController struct{}

func (userController UserController) Test(ctx *gin.Context) {
	user := models.User{
		RecordId:   nil,
		UserName:   "azzz",
		CreateTime: time.Now(),
	}
	err := models.CommonDao.Save(&user)
	if err != nil {
		common.ReturnResponse(ctx, 500, err.Error(), nil)
		return
	}
	common.ReturnResponse(ctx, 200, "ok", nil)
}
