package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 连接数据库
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		c.Ctx.WriteString("连接redis失败！")
		return
	}
	defer conn.Close()

	// 操作redis

	//// NO.1
	//conn.Send("set", "hello", "world")
	//conn.Flush()
	//rep,_ :=conn.Receive()
	//if err != nil {
	//	c.Ctx.WriteString("获取内容出错！")
	//	return
	//}
	//beego.Info(rep)
	//c.Ctx.WriteString("获取成功！")

	//// NO.2
	//rep, err := conn.Do("get", "hello")
	//res, err := redis.String(rep, err)
	//if err != nil {
	//	c.Ctx.WriteString("获取内容出错！")
	//	return
	//}
	//c.Ctx.WriteString(res)

	//// NO.3
	//rep, err := conn.Do("mget", "hello", "c1")
	//res, err := redis.Strings(rep, err)
	//if err != nil {
	//	c.Ctx.WriteString("获取内容出错！")
	//	return
	//}
	//beego.Info(res)
	//c.Ctx.WriteString("获取成功！")

	// NO.4
	rep, err := conn.Do("mget", "name", "age")
	res, err := redis.Values(rep, err)
	var name string
	var age int
	redis.Scan(res,&name,&age)
	if err != nil {
		c.Ctx.WriteString("获取内容出错！")
		return
	}
	beego.Info(name,age)
	c.Ctx.WriteString("获取成功！")
}
