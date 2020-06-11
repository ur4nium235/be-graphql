package schemas

import (
	"github.com/graphql-go/graphql"
	"server-graphql/internal/util"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 11/06/2020 23:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		util.HandleErrorPrintf(result.Errors, "")
	}
	return result
}