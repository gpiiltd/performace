// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"performance/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/lead",
			beego.NSInclude(
				&controllers.TeamLeadController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{},
			),
		),
		beego.NSNamespace("/team",
			beego.NSInclude(
				&controllers.TeamController{},
			),
		),
		beego.NSNamespace("/kpi",
			beego.NSInclude(
				&controllers.KPIController{},
			),
		),
		beego.NSNamespace("/objectives",
			beego.NSInclude(
				&controllers.ObjectiveController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TTController{},
			),
		),
		beego.NSNamespace("/report",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
