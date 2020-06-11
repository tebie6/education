package controllers

import (
	"education/models"
	"education/util"
)

// 声明一个AdminController控制器
type AdminController struct {
	baseController
}

// 主体展示 例如浏览器的是 GET 请求，那么默认就会执行 AdminController 下的 Get 方法
func (this *AdminController) Get() {

	// 调用展示权限列表
	permissionList := models.GetShowList()
	permissionListTree, _ := util.NodeDataMerge(permissionList,41)
	this.Data["navList"] = permissionListTree

	this.TplName = this.controllerName + "/main.html"
}

// console
func (this *AdminController) Console(){

	this.TplName = this.controllerName + "/console.html"
}

