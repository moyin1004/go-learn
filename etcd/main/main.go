package main

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const LockPrefix = "/lock/"

// DistributedLock 基于 etcd 的分布式锁
type DistributedLock struct {
	client *clientv3.Client
}

func NewDistributedLock(endpoints []string) (*DistributedLock, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	return &DistributedLock{client}, err
}

// Lock 尝试获取分布式锁（支持超时）
func (dl *DistributedLock) Lock(ctx context.Context, lockKey string) error {
	rsp, err := dl.client.Put(ctx, LockPrefix+lockKey, "1", clientv3.WithLease(10))
	if err != nil {
		return err
	}
	lock, err := dl.client.Get(ctx, LockPrefix)
	fmt.Println(rsp, lock)
	return nil
}

func (dl *DistributedLock) UnLock(ctx context.Context, lockKey string) error {
	rsp, err := dl.client.Delete(ctx, LockPrefix+lockKey)
	fmt.Println("Unlock", rsp)
	return err
}

// 使用示例
func main() {
	ctx := context.Background()
	endpoints := []string{"localhost:2379"} // etcd 地址
	dl, err := NewDistributedLock(endpoints)
	if err != nil {
		panic(err)
	}
	dl.Lock(ctx, "UUID1")
	dl.UnLock(ctx, "UUID1")
}
