package mysql

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slitproxy/user/app/model/mmysql"
	"slitproxy/user/pkg/logger"
	"slitproxy/user/pkg/mysql"
)

type User struct {
	c *gin.Context
	*mysql.DaoMysql
}

func NewUser(c *gin.Context) *User {
	return &User{
		DaoMysql: mysql.NewMysql(),
		c:        c,
	}
}

func (p *User) CreateUser(username, password string) (info *mmysql.UserDemo, err error) {
	orm := p.Orm(p.c)
	info = &mmysql.UserDemo{
		Username: username,
		Password: password,
	}
	// 执行插入
	sql := orm.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&info)
	})
	err = orm.Exec(sql).Error
	if err != nil {
		logger.Errorf(p.c, "CreateUser err : %v", err)
		return
	}
	return
}

func (p *User) GetUserByUsername(username string) (user *mmysql.UserDemo, err error) {
	orm := p.Orm(p.c)
	err = orm.Where(fmt.Sprintf("username = '%s'", username)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
		return
	}
	if err != nil {
		user = nil
		logger.Errorf(p.c, "GetUser err : %v", err)
		return
	}
	return
}

func (p *User) GetUserByUsernamePass(username, password string) (user *mmysql.UserDemo, err error) {
	orm := p.Orm(p.c)
	err = orm.Where(fmt.Sprintf("username = '%s' and password = '%s'", username, password)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
		return
	}
	if err != nil {
		user = nil
		logger.Errorf(p.c, "GetUserByUsernamePass err : %v", err)
		return
	}
	return
}
