package governance

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

func EtcdDial(cli *clientv3.Client, serviceName string) (*grpc.ClientConn, error) {
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(
		"etcd:///"+serviceName,
		grpc.WithResolvers(etcdResolver),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}
