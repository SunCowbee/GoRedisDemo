package main

import (
	"github.com/gitstliu/go-redis-cluster"
	"time"
	"fmt"
)

func main() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"192.168.150.20:7000", "192.168.150.20:7001", "192.168.150.20:7002", "192.168.150.21:7003", "192.168.150.21:7004", "192.168.150.21:7005"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		fmt.Println("redis集群连接失败！")
	}
	cluster.Do("set", "name", "roger")
	rep, err := cluster.Do("get", "name")
	name, _ := redis.String(rep, err)
	fmt.Println(name)
}
