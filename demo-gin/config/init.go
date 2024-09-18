package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetServerConfig() Server {
	return app.Server
}
func GetMysqlConfig() Mysql {
	return app.Database.Mysql
}
func GetRedisConfig() Redis {
	return app.Database.Redis
}

func GetEtcdConfig() Etcd {
	return app.Etcd
}

func GetMachineId() int {
	return app.Server.MachineId
}

var app Application

func InitConfig() {
	config, _ := ioutil.ReadFile("config/application.yml")
	err := yaml.Unmarshal(config, &app)
	if err != nil {
		panic(err)
	}
}

// Application -------------application----------
type Application struct {
	Server   Server
	Database Database
	Etcd     Etcd
}

// Server --------------server--------------
type Server struct {
	Mode      string `yaml:"mode"`
	Port      string `yaml:"port"`
	MachineId int    `yaml:"machineId"`
}

// Database ------------数据库配置-------------
type Database struct {
	Mysql Mysql
	Redis Redis
}

// Mysql ---------------mysql--------------
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

// Redis --------------redis--------------
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"Password"`
}

// Etcd
type Etcd struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	ServiceName string `yaml:"serivceName"`
	Version     string `yaml:"version"`
}
