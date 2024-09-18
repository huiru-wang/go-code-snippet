package governance

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
)

func register(service ServiceInstance, stop chan error) error {
	instance := NewServiceInstance()
	cli, err := ConfigEtcdClient(instance)
	if err != nil {
		return fmt.Errorf("service register failed:%v", err)
	}
	defer cli.Close()

	// 创建租约
	resp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		return fmt.Errorf("service register failed:%v", err)
	}
	leaseId := resp.ID
	// 注册服务：本质是向etcd添加一个kv
	manager, err := endpoints.NewManager(cli, instance.Name)
	if err != nil {
		return fmt.Errorf("service register failed:%v", err)
	}
	serviceInfo := instance.Name + "/" + instance.Addr
	err = manager.AddEndpoint(cli.Ctx(), serviceInfo, endpoints.Endpoint{Addr: instance.Addr}, clientv3.WithLease(leaseId))
	if err != nil {
		return fmt.Errorf("service register failed:%v", err)
	}
	// 心跳检测
	aliveChan, err := cli.KeepAlive(context.Background(), leaseId)
	if err != nil {
		return fmt.Errorf("service register failed:%v", err)
	}
	log.Printf("[%s] register service ok\n", instance.Addr)

	for {
		select {
		case err := <-stop:
			if err != nil {
				log.Println(err)
			}
			return err
		case <-cli.Ctx().Done():
			log.Println("service closed.")
			return nil
		case _, ok := <-aliveChan:
			if ok {
				log.Println("service closed for keep alive chan is closed.")
				return nil
			}
			log.Printf("Recv reply from service: %s/%s, ttl:%d", service, instance.Addr, resp.TTL)
		}
	}

}
