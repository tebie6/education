package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


type Filter struct {
	Page int
	PageSize int
}

func init(){

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"

	// 这个用来设置 driverName 对应的数据库类型
	 _ = orm.RegisterDriver("mysql", orm.DRMySQL)

	// 设置数据库连接参数
	_ = orm.RegisterDataBase("default", "mysql", dsn)

	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(AuthRole),
		new(AuthPermission),
		new(AuthRoleNode),
		new(Admin))
}


// 返回带有前缀的表名
func TableName(str string) string{

	return beego.AppConfig.String("dbprefix") + str
}
