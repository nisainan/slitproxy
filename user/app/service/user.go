package service

import (
	"github.com/gin-gonic/gin"
	"slitproxy/user/app/dao/mysql"
	"slitproxy/user/app/model/mmysql"
	"slitproxy/user/app/model/mparam"
	"slitproxy/user/pconst"
	"slitproxy/user/pkg/util"
)

func Register(c *gin.Context, param *mparam.Register) (code int, user *mmysql.UserDemo) {
	// 判断用户是否存在
	user, err := mysql.NewUser(c).GetUserByUsername(param.Username)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if user != nil && user.ID > 0 {
		return
	}
	// 执行注册功能
	user, err = mysql.NewUser(c).CreateUser(param.Username, util.NewMd5(param.Password))
	if err != nil {
		user = nil
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func Login(c *gin.Context, param *mparam.Login) (code int, user *mmysql.UserDemo) {
	// 判断用户是否存在
	user, err := mysql.NewUser(c).GetUserByUsernamePass(param.Username, util.NewMd5(param.Password))
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if user != nil && user.ID > 0 {
		return
	}
	code = pconst.CODE_COMMON_DATA_NOT_EXIST
	return
}
