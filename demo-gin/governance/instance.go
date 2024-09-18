package governance

import (
	"fmt"
	"go-gin/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServiceInstance struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Version string `json:"version"`
	Weight  int64  `json:"weight"`
}

func NewServiceInstance() *ServiceInstance {
	etcdConfig := config.GetEtcdConfig()
	return &ServiceInstance{
		Name:    etcdConfig.ServiceName,
		Addr:    fmt.Sprintf("%s:%s", etcdConfig.Host, etcdConfig.Port),
		Version: etcdConfig.Version,
	}
}

func ConfigEtcdClient(instance *ServiceInstance) (cli *clientv3.Client, err error) {
	clientConfig := clientv3.Config{
		Endpoints:   []string{instance.Addr},
		DialTimeout: time.Duration(5) * time.Second,
	}
	// 创建etcd client
	return clientv3.New(clientConfig)
}
