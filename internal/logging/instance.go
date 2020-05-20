package logging

/**
 *
 * @author: hoangtq
 * @timeCreate: 20/05/2020 23:12
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func Logging(pathToLog string) {
	config := Configuration{
		EnableConsole:     true,
		ConsoleLevel:      Info,
		ConsoleJSONFormat: true,
		EnableFile:        true,
		FileLevel:         Info,
		FileJSONFormat:    true,
		FileLocation:      pathToLog,
	}
	err := NewLogger(config, InstanceZapLogger)
	if err != nil {
		log.Fatalf("Can not init log %s", err.Error())
	}

	//logger := WithFields(Fields{"some-key": "some-value", "time": time.Now().Unix()})
	//logger.Debugf("Starting with zap (^.^)")
	//logger.Infof("I'm cool :))")

	//err = NewLogger(config, InstanceLogrusLogger)
	//if err != nil {
	//	log.Fatalf("Can not init log %s", err.Error())
	//}
	//logger = WithFields(Fields{"some-key": "some-value", "time": time.Now().Unix()})
	//logger.Debugf("Starting with logrus (^.^)")
	//
	//logger.Infof("I'm cool :))")
}