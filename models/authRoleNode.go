package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AuthRoleNode struct {
	Id 		int
	NodeId  int
	RoleId	int64
}

func (m *AuthRoleNode) TableName() string {
	return TableName("auth_role_node")
}

// 获取列表通过roleId
func (m *AuthRoleNode) GetListByRoleId(roleId int) ([] *AuthRoleNode) {

	o :=  orm.NewOrm()
	list := [] *AuthRoleNode{}

	o.QueryTable( new(AuthRoleNode).TableName() ).Filter("RoleId",roleId).All(&list)

	return list
}

// 清空数据关联节点 通过 roleId
func (m *AuthRoleNode) CleanAllByRoleId(RoleId int64) (error) {

	o := orm.NewOrm()
	_, err := o.QueryTable( new(AuthRoleNode) ).Filter("RoleId",RoleId).Delete()

	return err
}

// 批量插入关联节点
func (m *AuthRoleNode) CreateAuthNode(roleNode [] AuthRoleNode) (error){

	o := orm.NewOrm()
	_, err := o.InsertMulti(len(roleNode), roleNode)

	return err
}