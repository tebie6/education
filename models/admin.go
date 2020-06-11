package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id 		int64
	Username  string
	Password  string
	Status	int64
	RoleId 	int64
	RoleName string `orm:"-"`
	Phone   int64
	Email	string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Admin) TableName() string {
	return TableName("admin")
}

//type Filter struct {
//	Page int
//	PageSize int
//}

// 查询单条信息通过ID
func (m *Admin) GetInfoById(id int64) (Admin, string){

	o := orm.NewOrm()
	info := Admin{Id:id}

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
func (m *Admin) GetListByWhere(filter *Filter) ([] *Admin, int64){

	o :=  orm.NewOrm()
	list := make([] *Admin, 0)

	if filter.Page < 1{
		filter.Page = 1
	}

	// 接受页码
	if filter.Page < 1 {
		filter.Page = 1
	}

	// 计算偏移量
	offset := (filter.Page - 1) * filter.PageSize

	query := o.QueryTable( new(Admin).TableName() )

	count, _ := query.Count()

	if count > 0 {

		// 查询角色
		rolelist := make(map[int64]string,0)
		roleListRes := make([] *AuthRole,0)
		o.QueryTable( new(AuthRole).TableName() ).All(&roleListRes)

		for _, v := range roleListRes {
			rolelist[v.Id] = v.RoleName
		}

		query.OrderBy("-created_at").Limit(filter.PageSize, offset).All(&list)

		for _, v := range list {
			v.RoleName = rolelist[v.RoleId]
		}
	}

	beego.Error(list)
	return list, count
}