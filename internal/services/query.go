package services

import (
	"errors"
	"server-graphql/internal/util"
	"strconv"
	"strings"
	"sync"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 13/06/2020 17:23
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (gs *GraphQLServices) getPostList(args map[string]interface{}) (interface{}, error) {
	var postIDList = make([]string, 0)
	valid := make(map[string]bool)
	if val, ok := args["ids"]; ok && len(val.(string)) > 0 {
		tmp := strings.Split(val.(string), ",")
		for index := range tmp {
			if _, ok := valid[tmp[index]]; !ok && len(tmp[index]) > 0 {
				postIDList = append(postIDList, tmp[index])
				valid[tmp[index]] = true
			}
		}
	}
	if len(postIDList) == 0 {
		return make([]interface{}, 0), errors.New("ids empty")
	}
	if len(postIDList) > maxItemHandle {
		return make([]interface{}, 0), errors.New("Too Many Items, Max values is " + strconv.Itoa(maxItemHandle))
	}
	wait := &sync.WaitGroup{}
	var result interface{}
	wait.Add(1)
	go func() {
		defer func() {
			wait.Done()
			if err := recover(); err != nil {
				util.HandleErrorPrintf(err, "")
			}

		}()
		result = gs.processGetUser(postIDList...)
	}()
	wait.Wait()
	return result, nil
}
