package controllers

import (
	"github.com/astaxie/beego"
	"education/models"
	"strconv"
	"strings"
	"time"
)

type AuthController struct {
	baseController
}

// 定义变量
var (
	AuthRoleModel *models.AuthRole
	AdminModel *models.Admin
	AuthPermissionModel *models.AuthPermission
	AuthRoleNodeModel *models.AuthRoleNode
)

// 角色编辑页面
func (this *AuthController) Role(){

	id, _ := this.GetInt64("id")

	if id != 0 {

		role, _ := AuthRoleModel.GetInfoById(id)
		this.Data["role"] = role
	}

	this.TplName = this.controllerName + "/roleform.html"
}

// 角色保存
func (this *AuthController) RoleSave(){

	authids := this.ContactPost("authids")

	beego.Debug(authids)
	beego.Debug(authids["0"])

	id, _ := this.GetInt64("id")

	role := models.AuthRole{}
	role.RoleName = this.GetString("role_name")
	role.RoleAliasName = this.GetString("role_alias_name")
	role.Descr = this.GetString("descr")
	role.UpdatedAt = time.Now()

	if id == 0 {

		role.CreatedAt = time.Now()
		if id, err := AuthRoleModel.Create(&role); err != nil {

			this.renderJson(500, "角色保存失败"+err.Error(), nil)
		} else {

			// 清空旧数据关联节点 由于ORM 需要主键 所以不能使用 orm.Delete
			err := AuthRoleNodeModel.CleanAllByRoleId(id)
			if err != nil {
				beego.Error( err.Error() )
			}

			// 批量插入关联节点
			//roleNode := [] models.AuthRoleNode{
			//	{NodeId:1, RoleId:id},
			//	{NodeId:3, RoleId:id},
			//	{NodeId:5, RoleId:id},
			//}
			if len(authids) > 0 {

				roleNode := [] models.AuthRoleNode{}

				// 切片 嵌套 map
				for _, _v := range authids {
					// sting 转换 int
					_v, _ := strconv.Atoi(_v)

					//string 转换 int64
					//_v, _ := strconv.ParseInt(_v, 10, 64)
					roleNode = append( roleNode, models.AuthRoleNode{NodeId: _v, RoleId:id})
				}
				err := AuthRoleNodeModel.CreateAuthNode(roleNode)
				if err != nil {
					beego.Error( err.Error() )
				}
			}

			this.renderJson(200, "角色保存成功", map[string] string{"url":"/auth/roleList"})
		}

	} else {

		role.Id = id
		_, err := AuthRoleModel.Update(&role)
		if err != nil {

			this.renderJson(500, "角色保存失败"+err.Error(), nil)
		} else {

			// 清空旧数据关联节点 由于ORM 需要主键 所以不能使用 orm.Delete
			err := AuthRoleNodeModel.CleanAllByRoleId(id)
			if err != nil {
				beego.Error( err.Error() )
			}

			// 批量插入关联节点
			//roleNode := [] models.AuthRoleNode{
			//	{NodeId:1, RoleId:id},
			//	{NodeId:3, RoleId:id},
			//	{NodeId:5, RoleId:id},
			//}
			if len(authids) > 0 {

				roleNode := [] models.AuthRoleNode{}

				// 切片 嵌套 map
				for _, _v := range authids {
					// sting 转换 int
					_v, _ := strconv.Atoi(_v)
					roleNode = append( roleNode, models.AuthRoleNode{NodeId: _v, RoleId:id})
				}
				err := AuthRoleNodeModel.CreateAuthNode(roleNode)
				if err != nil {
					beego.Error( err.Error() )
				}
			}

			this.renderJson(200, "角色保存成功", map[string] string{"url":"/auth/roleList"})
		}
	}

	//this.ServeJSON()
}

// 角色列表
func (this *AuthController) RoleList(){

	if this.Ctx.Request.Method == "POST" {

		// 接受设置参数
		filter := models.Filter{}
		filter.Page, _ = this.GetInt("page")
		filter.PageSize, _ = this.GetInt("limit")

		// 调用 GetAll 方法
		roleList, count := AuthRoleModel.GetAll(&filter)

		// 设置扩展参数
		this.setInputData(Response{Count:count})

		// 渲染并输出数据
		this.renderJson(0, "success", roleList)
	}

	this.TplName = this.controllerName + "/role_list.html"
}


// 删除角色
func (this *AuthController) RoleDelete(){

	id, _ := this.GetInt64("id")

	if id == 0 {

		// TODO 此处需要优化
		ids := strings.Join(this.GetStrings("ids[]"), ",")

		res := AuthRoleModel.BatchDelete(ids)
		if res == false {
			this.renderJson(500, "删除失败", nil)

		} else {
			this.renderJson(200, "删除成功", nil)
		}

	} else {

		if _, err := AuthRoleModel.Delete(id); err != nil {
			this.renderJson(500, "删除失败", nil)

		} else {
			this.renderJson(200, "删除成功", nil)
		}
	}

}

// 权限列表
func (this *AuthController) PermissionList(){

	this.TplName = this.controllerName + "/permission_list.html"
}

// 获取列表数据 因为treetable 不支持post请求 所以单起接口
func (this *AuthController) PermissionListApi() {


	act := this.GetString("act")

	if act == "authtree" {

//		json := `
//{
//  "code": 0,
//  "msg": "获取成功",
//  "data": {
//    "list": [
//      { "id": 1, "name": "用户管理", "pid": 0 },
//      { "id": 2, "name": "用户组管理", "pid": 0 },
//      { "id": 3, "name": "角色管理", "pid": 2 },
//      { "id": 4, "name": "添加角色", "pid":  3},
//      { "id": 5, "name": "角色列表", "pid": 3 },
//      { "id": 6, "name": "管理员管理", "pid": 0 },
//      { "id": 7, "name": "添加管理员", "pid": 6 },
//      { "id": 8, "name": "管理员列表", "pid": 6 }
//    ],
//    "checkedId": [ 1,  2, 3, 4 ]
//  }
//}`

		roleId, _ := this.GetInt("role_id")

		// 查询权限节点
		list := AuthPermissionModel.GetAll()

		// 查询当前角色权限节点
		checkedId := [] int{}
		if roleId !=0 {

			roleNode := AuthRoleNodeModel.GetListByRoleId(roleId)

			// TODO 此处应该有更好的方法实现
			for _, _v := range roleNode {
				checkedId = append(checkedId, _v.NodeId)
			}
		}

		// 整合返回数据
		returnList := map[string] interface{}{
			"list" : list,
			"checkedId" : checkedId,
		}

		// 设置扩展参数
		this.setInputData(Response{Count:0})

		// 渲染并输出数据
		this.renderJson(0, "success", returnList)

		//this.Ctx.WriteString(json)
	} else {
		list := AuthPermissionModel.GetAll()

		// 设置扩展参数
		this.setInputData(Response{Count:0})

		// 渲染并输出数据
		this.renderJson(0, "success", list)
	}

}

// 权限编辑页
func (this *AuthController) Permission() {

	if this.Ctx.Request.Method == "POST" {
		act := this.GetString("act")

		// 获取类型数据
		if act == "getSelectTypeData" {

			typeData := `[
				{
					"id" : "1",
					"name" : "顶级应用"
				},
				{
					"id" : "2",
					"name" : "模块"
				},
				{
					"id" : "3",
					"name" : "控制器/方法"
				}
			]`

			this.Ctx.WriteString(typeData)

		} else if act == "getParentSelectData" {

			level, _ := this.GetInt("level")
			level = level - 1

			list := AuthPermissionModel.GetListByLevel(level)

			this.renderJson(200, "success", list)
		}


	}

	this.TplName = this.controllerName + "/permission.html"
}

// 权限保存
func (this *AuthController) PermissionSave(){

	id, _ := this.GetInt("id")

	permission := models.AuthPermission{}
	permission.Title = this.GetString("title")
	permission.Pid, _ = this.GetInt("pid")
	permission.Level, _ = this.GetInt("level")
	permission.Status, _ = this.GetInt("status")
	permission.IsShow, _ = this.GetInt("is_show")
	permission.Route = this.GetString("route")
	permission.IsDel = 0
	permission.UpdatedAt = time.Now()

	if id == 0 {

		permission.CreatedAt = time.Now()
		if _, err := this.o.Insert(&permission); err != nil {
			this.renderJson(500, "创建失败" + err.Error(), nil)
		} else {
			this.renderJson(200, "创建成功", nil)
		}

	} else {

		permission.Id = id

		if _, err := this.o.Update(&permission, "Title","Pid","Pid","Level","Status","IsShow","Route","UpdatedAt"); err != nil {
			this.renderJson(500, "更新失败 " + err.Error(), nil)
		} else {
			this.renderJson(200, "更新成功", nil)
		}
	}
}

// 权限删除
func (this *AuthController) PermissionDel(){

	id, _ := this.GetInt("id")

	if id == 0 {
		this.renderJson(500, "参数错误", nil)
	}

	info := models.AuthPermission{}

	info.Id = id
	info.IsDel = 1		//TODO 此处应该使用常量
	info.UpdatedAt = time.Now()

	if _, err := this.o.Update(&info, "IsDel","UpdatedAt"); err != nil {
		this.renderJson(500, "删除失败 " + err.Error(), nil)
	} else {
		this.renderJson(200, "删除成功", nil)
	}

}


// 后台用户
func (this *AuthController) AdminList(){

	if this.Ctx.Request.Method == "POST" {

		// 接受设置参数
		filter := models.Filter{}
		filter.Page, _ = this.GetInt("page")
		filter.PageSize, _ = this.GetInt("limit")

		list, count := AdminModel.GetListByWhere(&filter)

		this.setInputData(Response{Count:count})

		this.renderJson(0, "success", list)
	}

	roleList, _ := AuthRoleModel.GetAll(&models.Filter{})

	this.Data["roleList"] = roleList

	this.TplName = this.controllerName + "/admin_list.html"
}

// 后台用户编辑
func (this *AuthController) Admin(){

	id, _ := this.GetInt64("id")

	if id != 0 {
		info, _ := AdminModel.GetInfoById(id)
		this.Data["info"] = info
	}

	this.TplName = this.controllerName + "/admin.html"
}

// 后台用户保存
func (this *AuthController) AdminSave(){

	id, _ := this.GetInt64("id")

	info := models.Admin{}
	info.Username = this.GetString("username")
	info.Phone, _ = this.GetInt64("phone")
	info.Email = this.GetString("email")
	info.RoleId, _ = this.GetInt64("role")
	info.Status, _ = this.GetInt64("status")
	info.UpdatedAt = time.Now()

	if id == 0 {

		info.CreatedAt = time.Now()
		if _, err := this.o.Insert(&info); err != nil {

			this.renderJson(500, "创建失败 "+err.Error(),nil)
		} else {

			this.renderJson(200, "创建成功",nil)
		}
	} else {

		info.Id = id

		if _, err := this.o.Update(&info,"Username","Phone","Email","RoleId","Status","UpdatedAt"); err != nil {

			this.renderJson(500, "更新失败 "+err.Error(), nil)
		} else {

			this.renderJson(200, "更新成功", nil)
		}
	}
}