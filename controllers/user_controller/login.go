package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	Username string `json:"username" form:"username" validate:"gt=4,lt=20,required"`
	Password string `form:"password" validate:"required,gt=4,lt=20"`
}

// 登陆逻辑
// 校验账号 -> 校验密码 -> 校验黑名单 -> 消息队列(更新登陆时间) -> res
func Login(ctx *gin.Context) {
	errRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	var loginInfo LoginInfo
	if err := helper.BindWithValid(ctx, &loginInfo); err != nil {
		errRes.Err = err
		errRes.Send(ctx)
		return
	}

	// 检查用户是否存在
	user, err := models.FindOneByUsername(loginInfo.Username)

	if err != nil {
		errRes.Err = robust.CANNOT_FIND_USER
		errRes.Send(ctx)
		return
	}

	if helper.IsEmpty(user) || user.ID == "" {
		errRes.Err = robust.LOGIN_USER_DOES_NOT_EXIST
		errRes.Send(ctx)
		return
	}

	// 密码错误
	if utils.Hash(loginInfo.Password) != user.Password {
		errRes.Err = robust.LOGIN_PASSWORD_NOT_VALID
		errRes.Send(ctx)
		return
	}

	// 验证是否在黑名单
	//if middlewares.IsExistInBlackSet(user.ID) {
	//	errRes.Err = robust.JWT_NOT_ALLOWED
	//	errRes.Send(ctx)
	//	return
	//}

	go user.UpdateLoginTime()

	claims := utils.Claims{
		UserId: user.ID,
	}

	res.Data = gin.H{
		"jwt":      claims.SignJWT(),
		"userInfo": user,
	}
	res.Send(ctx)
}
