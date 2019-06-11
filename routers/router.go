package routers

import (
	"ServerUtils/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainUtilController{})
	beego.Include(&controllers.SSHController{})
}
