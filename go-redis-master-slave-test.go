package main

//import (
//	"github.com/gomodule/redigo/redis"
//	"fmt"
//)
//
//func main() {
//	// 连接master
//	master, err := redis.DialURL("redis://192.168.150.20:6379")
//	if err != nil {
//		fmt.Println("连接master失败！")
//	}
//	defer master.Close()
//	// 设置数据
//	master.Do("set", "d1", "hello world")
//
//	// 连接slave
//	slave, err := redis.DialURL("redis://192.168.150.20:6378")
//	if err != nil {
//		fmt.Println("连接slave失败！")
//	}
//	defer slave.Close()
//	// 获取数据
//	rep, err := slave.Do("get", "d1")
//	res, err := redis.String(rep, err)
//	fmt.Println(res)
//}
