package server

import (
	"github.com/urfave/cli"
	"slitproxy/user/pkg/confer"
	"slitproxy/user/pkg/logger"
	"slitproxy/user/pkg/mysql"
)

func InitService(c *cli.Context) (err error) {
	if err = confer.Init(c.String("c")); err != nil {
		return
	}
	cfg := confer.GlobalConfig()
	logger.Init(&logger.Config{
		Level:       logger.LogLevel(),
		Filename:    logger.LogFile(),
		SendToFile:  logger.SendLogToFile(),
		Development: confer.ConfigEnvIsDev(),
	})
	if err = mysql.Init(&cfg.Mysql); err != nil {
		logger.Errorf(nil, "mysql init error : %v", err)
		return
	}
	return
}
