package services

import (
	"server-graphql/internal/logging"
	"server-graphql/internal/micro"
	"server-graphql/internal/schemas"
	"server-graphql/internal/util"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/06/2020 05:24
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func InitGraphQLServices() (*GraphQLServices, error) {
	gs := &GraphQLServices{}

	err := gs.initConf(pathConf)

	if err != nil {
		return nil, err
	}

	logging.Logging(gs.conf.LogRequest)

	err = gs.initSchemas()

	if err != nil {
		return nil, err
	}
	gs.initMicroServices()
	gs.initRouter()


	return gs, nil
}

func (sp *GraphQLServices) ListenAndServe() {
	if sp.conf.ModeDebug == 0 {
		util.HandlePrintf("Listening and serving HTTP on " + sp.conf.ServerAddr)
	}
	err := sp.router.Run(sp.conf.ServerAddr)
	if err != nil {
		util.HandleFatalf(err)
	}
}

func (gs *GraphQLServices) initConf(pConf string) error {
	gs.conf = &Config{}
	return util.LoadFileConfig(pConf, gs.conf)
}

func (gs *GraphQLServices) initSchemas() error {
	//Post
	//queries := make(map[string]schemas.HandlerFuncQuery)
	//queries["post_list"] = gs.getPostList
	gs.postQuery = &schemas.PostQuery{}

	gs.postQuery.RegistryQueryFunc(nil, gs.getPostList)
	gs.postQuery.RegistryBatchFunc(gs.getUserInfoFunc)

	return gs.postQuery.InitSchema()
}

func (gs *GraphQLServices) initMicroServices()  {
	util.HandlePrintf("Init Microservices: ")
	gs.userManager = micro.InitServiceUser(gs.conf.UserInfoAPI)

}

