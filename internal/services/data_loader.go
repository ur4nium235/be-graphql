package services

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"server-graphql/internal/bjson"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 11/06/2020 22:52
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (gs GraphQLServices) getUserInfoFunc(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	tmp := gs.userManager.GetUserInfo(context.TODO(), keys.Keys()...)

	for _, obj := range keys {
		user := bjson.UserInfo{}

		if val, ok := tmp[obj.String()]; ok {
			user.UserID = val.UserID
			user.Username = val.Username
			user.FullName = val.FullName
			user.Address = val.Address
			user.City = val.City
			user.Type = val.Type
			user.Role = val.Role
		}
		results = append(results, &dataloader.Result{Data:user})
	}
	return results
}
