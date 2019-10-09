package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "AddNewMember",
            Router: `/member/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TokenController"] = append(beego.GlobalControllerRouter["performance/controllers:TokenController"],
        beego.ControllerComments{
            Method: "GetTokenString",
            Router: `/token/:email`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:UserController"] = append(beego.GlobalControllerRouter["performance/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateProfile",
            Router: `/update/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
