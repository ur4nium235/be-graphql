package services

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server-graphql/internal/schemas"
	"server-graphql/internal/util"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 11/06/2020 23:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (gs *GraphQLServices) postHandlePostQuery(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			util.HandlePrintf(err)
		}
	}()
	body := ctx.Request.Body
	data, _ := ioutil.ReadAll(body)
	if len(string(data)) > 0 {
		result := schemas.ExecuteQuery(string(data), gs.postQuery.GetSchema())
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.String(http.StatusBadRequest, "Please use GraphQL format.")
	}
}
