package test

import (
	config2 "be-graphql/internal/config"
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/05/2020 23:15
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func TestLoadFileConfig(t *testing.T)  {
	var config config2.TestConfig
	if _, err := toml.DecodeFile("/storage/data/Work/MyRepositories/be-graphql/configs/config_example.conf", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
		config.Owner.Name, config.Owner.Org, config.Owner.Bio,
		config.Owner.DOB)
	fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
		config.DB.Server, config.DB.Ports, config.DB.ConnMax,
		config.DB.Enabled)
	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}
	fmt.Printf("Client data: %v\n", config.Clients.Data)
	fmt.Printf("Client hosts: %v\n", config.Clients.Hosts)
}