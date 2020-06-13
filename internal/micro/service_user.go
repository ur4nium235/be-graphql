package micro

import (
	"context"
	"fmt"
	"server-graphql/internal/bjson"
	"server-graphql/internal/util"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 11/06/2020 22:55
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func InitServiceUser(api string) *ServiceUser {
	util.HandlePrintf("- Microservice ServiceUser")
	var ca =&ServiceUser{}

	ca.client = util.CreateFastClient(100);
	ca.api = api
	return ca
}

func (ca *ServiceUser) GetUserInfo(_ context.Context, userIDs ...string) map[string]bjson.UserInfo  {

	fmt.Println("Get UserInfo")

	if len(userIDs) == 0 {
		return make(map[string]bjson.UserInfo)
	}

	// Todo call api or connect db get user info

	// fix cứng (^_^)

	var out = make(map[string]bjson.UserInfo)

	user := bjson.UserInfo{
		UserID:   "123456789",
		Username: "ur4nium235",
		FullName: "Uranium235",
		Address:  "HP",
		City:     "HN",
		//Type:     1,
		Role:     "admin",
	}
	out[user.UserID] = user
	return out
}


