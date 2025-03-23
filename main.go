package main

import "demo/util/mysql"

func main() {

	// 数据库初始化
	mysql.Init()

	// 路由配置
	routeInit()

}
