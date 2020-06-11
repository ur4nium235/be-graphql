package micro

import (
	"github.com/valyala/fasthttp"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/06/2020 05:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type ServiceUser struct {
	client *fasthttp.Client
	api    string
}