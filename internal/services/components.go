package services

import (
	"server-graphql/internal/micro"
	"server-graphql/internal/schemas"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/06/2020 05:25
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type GraphQLServices struct {
	router *gin.Engine
	conf      *Config

	logServer *lumberjack.Logger
	postQuery *schemas.PostQuery

	//micro
	userManager    *micro.ServiceUser
}

type Config struct {
	//service
	ServerAddr string
	ModeDebug  int
	ModeBeta   int //1 cho bản beta. 0 cho prod
	//log
	LogRequest string

	UserInfoAPI   string
}
