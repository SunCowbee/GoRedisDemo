# GoRedisDemo

ENVIRONMENT REQUIRED:
#ubuntu
#go
#beego
$ go get -u -v github.com/astaxie/beego
$ go get -u -v github.com/beego/bee
##配置beego执行文件环境变量：$GOPATH/bin
	$ vim .bashrc
	//在最后一行插入
	export PATH="$GOPATH/bin:$PATH"
	//然后保存退出
	$ source .bashrc
	
===go-redis-test.go===	
#redis
cd /etc/redis/
sudo redis-server ./redis.conf 
redis-cli
#go-redis-api
go get -u -v github.com/gomodule/redigo/redis
#go-redis-api-doc
https://godoc.org/github.com/gomodule/redigo/redis

===go-redis-master-slave-test.go===	
#redis
cd /etc/redis/
sudo redis-server ./master.conf 
sudo redis-server ./slave.conf 
redis-cli -h 192.168.150.20 -p 6379
redis-cli -h 192.168.150.20 -p 6378
#go-redis-api
go get -u -v github.com/gomodule/redigo/redis
#go-redis-api-doc
https://godoc.org/github.com/gomodule/redigo/redis

===go-redis-cluster-test.go===	
#redis
cd /etc/redis/
*.conf
	port 7000
	bind 192.168.110.37
	daemonize yes
	pidfile 7000.pid
	cluster-enabled yes
	cluster-config-file 7000_node.conf
	cluster-node-timeout 15000
	appendonly yes
192.168.150.20 
sudo redis-server ./7000.conf 
sudo redis-server ./7001.conf 
sudo redis-server ./7002.conf 
192.168.150.21
sudo redis-server ./7003.conf 
sudo redis-server ./7004.conf 
sudo redis-server ./7005.conf 
sudo cp /usr/share/doc/redis-tools/examples/redis-trib.rb /usr/local/bin/
#ruby
sudo apt-get install ruby
gem sources --add https://gems.ruby-china.com --remove https://rubygems.org/
// 通过 gem 安装 redis 的相关依赖
sudo gem install redis
创建集群
redis-trib.rb create --replicas 1 192.168.150.20:7000 192.168.150.20:7001 192.168.150.20:7002 192.168.150.21:7003 192.168.150.21:7004 192.168.150.21:7005
进入redis
redis-cli -h 192.168.150.20 -c -p 7000
#go-redis-cluster-api
go get -u -v github.com/gitstliu/go-redis-cluster
go get -u -v github.com/gomodule/redigo/redis
#go-redis-cluster-api-doc
https://godoc.org/github.com/gomodule/redigo/redis
http://www.cnblogs.com/wuxl360/p/5920330.html
https://github.com/gitstliu/go-redis-cluster
