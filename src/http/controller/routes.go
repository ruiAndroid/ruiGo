package controller


func RegisterRoutes(){
	//注册首页相关路由
	//new(IndexController).RegistRoute()
	//注册收集用户bug路由
	new(UserBugTrackerController).RegistRoute()
}
