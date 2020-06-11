package controllers

type MainController struct {
	baseController
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = c.controllerName + "/main.html"
}
