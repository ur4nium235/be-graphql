package services

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-graphql/internal/logging"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 11/06/2020 23:23
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (gs *GraphQLServices) initRouter() {
	if gs.conf.ModeDebug != debug {
		gin.SetMode(gin.ReleaseMode)
	}

	gs.router = gin.New()
	gs.router.Use(logging.LoggerWithWriter(gs.logServer))
	gs.router.Use(gin.Recovery())
	gs.router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiV1 := gs.router.Group("/api/v1")
	apiV1.Handle(http.MethodPost, "/post/query", gs.postHandlePostQuery)
}