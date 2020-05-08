package util

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"runtime"
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
		log.Printf("[I] %v", msg)
	}
}

func HandleErrorPrintf(err interface{}, msg string) {
	if err != nil {
		_, fn, line, ok := runtime.Caller(1)
		if ok {
			if len(msg) > 0 {
				log.Printf("[E] %v %s:%d \t %v", err, fn, line, msg)
			} else {
				log.Printf("[E] %v %s:%d", err, fn, line)
			}
		} else if len(msg) > 0 {
			log.Printf("[E] %v", msg)
		}
	}
}

func HandleFatalf(err interface{}) {
	if err != nil {
		_, fn, line, ok := runtime.Caller(1)
		if ok {
			log.Printf("[E] %v %s:%d", err, fn, line)
		}
	}
}
