package schemas

import (
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/06/2020 04:42
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type HandlerFuncQuery func(args map[string]interface{}) (interface{}, error)

type PostQuery struct {
	schema graphql.Schema
	//Func For Query
	getPostFunc     HandlerFuncQuery
	getPostListFunc HandlerFuncQuery
	//DataLoader
	userPostLoader    *dataloader.Loader
}



