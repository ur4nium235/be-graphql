package config

import "time"

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/05/2020 23:20
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type TestConfig struct {
	Title   string
	Owner   OwnerInfo
	DB      Database `toml:"database"`
	Servers map[string]Server
	Clients Clients
}

type OwnerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type Database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type Server struct {
	IP string
	DC string
}

type Clients struct {
	Data  [][]interface{}
	Hosts []string
}