package services

import (
	"context"
	"strings"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 13/06/2020 17:26
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (gs *GraphQLServices) processGetUser(ids ...string) interface{} {
	out := make([]interface{}, 0)
	posts := make([]string, 0)
	for _, id := range ids {
		tmp := strings.Split(id, "_")
		posts = append(posts, tmp[0])
	}

	data := gs.userManager.GetUserInfo(context.TODO(), posts...)

	for _, id := range ids {
		if obj, ok := data[id]; ok {
			obj.UserID = id
			out = append(out, obj)
		}
	}
	return out
}
