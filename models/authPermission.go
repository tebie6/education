package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"

)

type AuthPermission struct {
	Id int
	Title string
	Pid int
	Level int
	Status int
	Route string
	IsShow int
	IsDel int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 用于处理树结构
type AuthPermissionTree struct {
	Id int
	Title string
	Pid int
	Level int
	Status int
	Route string
	IsShow int
	IsDel int
	CreatedAt time.Time
	UpdatedAt time.Time
	Child [] *AuthPermissionTree
}

const (
	// 展示状态
	IS_SHOW_T = 1
	IS_SHOW_F = 0
)

func (m *AuthPermission) TableName() string {
	return TableName("auth_permission")
}

// 获取全部数据
func (m *AuthPermission) GetAll() ([] *AuthPermission) {

	o :=  orm.NewOrm()
	list := [] *AuthPermission{}

	// 此处可以做 缓存优化处理
	o.QueryTable( new(AuthPermission).TableName() ).Filter("IsDel",0).All(&list)

	return list
}

// 获取展示权限列表
func GetShowList() ([] *AuthPermission)  {

	o :=  orm.NewOrm()
	list := [] *AuthPermission{}

	// 此处可以做 缓存优化处理
	o.QueryTable( new(AuthPermission).TableName() ).Filter("IsShow",IS_SHOW_T).All(&list)

	return list
}

// 获取列表通过roleId
func (m *AuthPermission) GetListByLevel(level int) ([] *AuthPermission) {

	o :=  orm.NewOrm()
	list := [] *AuthPermission{}

	o.QueryTable( new(AuthPermission).TableName() ).Filter("Level",level).All(&list)

	return list
}
