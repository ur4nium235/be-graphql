package util

import (
	"github.com/valyala/fasthttp"
	"server-graphql/internal/logging"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/05/2020 21:37
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func LoadFileConfig(fileName string, config interface{}) error {
	// lấy metadata (kích thước, thời gian tạo, etc.) về một tập tin mà không cần mở nó
	_, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	// độc nội dung của fileName và giải mã nó
	if _, err := toml.DecodeFile(fileName, config); err != nil {
		return err
	}
	return err
}

func HandlePrintf(msg interface{}) {
	if msg != nil {
		contextLogger := logging.WithFields(logging.Fields{"time":time.Now().Unix()})
		contextLogger.Infof("%v", msg)
	}
}

func HandleErrorPrintf(err interface{}, msg string) {
	if err != nil {
		logger := logging.WithFields(logging.Fields{"time": time.Now().Unix()})
		logger.Errorf("%v", err)
	}
}

func HandleFatalf(err interface{}) {
	if err != nil {
		logger := logging.WithFields(logging.Fields{"time": time.Now().Unix()})
		logger.Fatalf("%v", err)
	}
}

func CreateFastClient(maxConnsPerHost int) *fasthttp.Client {
	return &fasthttp.Client{MaxConnsPerHost: maxConnsPerHost, MaxIdleConnDuration: 5 * time.Second,
		ReadTimeout: 2 * time.Second, WriteTimeout: 2 * time.Second,}
}

