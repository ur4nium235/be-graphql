package test

import (
	config2 "be-graphql/internal/config"
	"be-graphql/internal/logging"
	"be-graphql/internal/util"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"testing"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/05/2020 23:15
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func init() {
	logging.Logging("/storage/data/Work/MyRepositories/be-graphql/logs/test.logs")
}

func TestMain(m *testing.M)  {
	//setup()
	code := m.Run()
	os.Exit(code)
}


func TestLoadFileConfig(t *testing.T)  {
	var config config2.TestConfig
	if _, err := toml.DecodeFile("/storage/data/Work/MyRepositories/be-graphql/configs/config_example.conf", &config); err != nil {
		util.HandlePrintf(err)
		return
	}

	util.HandlePrintf(fmt.Sprintf("Title: %v", config.Title))
	util.HandlePrintf(fmt.Sprintf("Owner: %s (%s, %s), Born: %s",
	config.Owner.Name, config.Owner.Org, config.Owner.Bio,
		config.Owner.DOB))
	util.HandlePrintf(fmt.Sprintf("Database: %s %v (Max conn. %d), Enabled? %v",
		config.DB.Server, config.DB.Ports, config.DB.ConnMax,
		config.DB.Enabled))


	for serverName, server := range config.Servers {
		util.HandlePrintf(fmt.Sprintf("Server: %s (%s, %s)", serverName, server.IP, server.DC))
	}
	util.HandlePrintf(fmt.Sprintf("Client data: %v", config.Clients.Data))
	util.HandlePrintf(fmt.Sprintf("Client hosts: %v", config.Clients.Hosts))
}