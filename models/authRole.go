package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AuthRole struct {
	Id int64
	RoleName string
	RoleAliasName string
	Status int
	Descr  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *AuthRole) TableName() string {
	return TableName("auth_role")
}

// 查询单条信息通过ID
func (m *AuthRole) GetInfoById(id int64) (AuthRole, string){

	o := orm.NewOrm()
	info := AuthRole{Id:id}

	err := o.Read(&info)
	msg := "success"

	if err == orm.ErrNoRows {
		msg = "查询不到"
	} else if err == orm.ErrMissPK {
		msg = "找不到主键"
	}

	return info, msg
}

// 查询所有
func (m *AuthRole) GetAll(filter *Filter) ([] *AuthRole, int64){

	o :=  orm.NewOrm()
	list := [] *AuthRole{}

	if filter.Page < 1{
		filter.Page = 1
	}

	// 接受页码
	if filter.Page < 1 {
		filter.Page = 1
	}

	// 计算偏移量
	offset := (filter.Page - 1) * filter.PageSize

	query := o.QueryTable( new(AuthRole).TableName())

	count, _ := query.Count()

	if count > 0 {

		query.OrderBy("-created_at").Limit(filter.PageSize, offset).All(&list)
	}


	return list, count
}

// 批量删除
func (m *AuthRole) BatchDelete(ids string) bool {

	beego.Debug(ids)
	o := orm.NewOrm()
	//_, err := o.Raw("UPDATE `" + new(AuthRole).TableName() + "` SET status = ? WHERE id IN(?)", 1, ids).Exec()
	//_, err := o.Raw("DELETE FROM `" + new(AuthRole).TableName() + "` WHERE id IN(?)", ids).Exec()


	// TODO 由于 ids 为 1,2,3 格式的字符串 每个逗号匹配一个"？"问号 所以只能操作第一个值 例如 where id in(1)
	//_, err := o.Raw("DELETE FROM `" + new(AuthRole).TableName() + "` WHERE id IN(?)", ids).Exec()
	_, err := o.Raw("DELETE FROM `" + new(AuthRole).TableName() + "` WHERE id IN(" + ids + ")").Exec()
	if err != nil {
		beego.Debug(err.Error())
		return false
	}

	return true
}

// 创建角色
func (m *AuthRole) Create(role *AuthRole) (int64, error){

	o := orm.NewOrm()
	id, err := o.Insert(role)

	return id, err
}

// 更新角色
func (m *AuthRole) Update(role *AuthRole) (int64, error){

	o := orm.NewOrm()
	id, err := o.Update(role, "RoleName", "RoleAliasName", "Descr", "UpdatedAt")

	return id, err
}

// 删除角色
func (m *AuthRole) Delete(id int64) (int64, error){

	o := orm.NewOrm()
	id, err := o.Delete(&AuthRole{Id:id})

	return id, err
}