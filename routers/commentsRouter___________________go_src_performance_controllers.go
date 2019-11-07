package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "AssignKPI",
            Router: `/assign/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "AcceptTeamInvitation",
            Router: `/accept/:teamid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "GetMyPendingTeam",
            Router: `/invitations/pending`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "GetMyTeamInformation",
            Router: `/myteam`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "GetTeamReport",
            Router: `/report`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "AddNewMember",
            Router: `/member/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "GetMyTeamInfo",
            Router: `/myteam/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "GetMyPendingTeam",
            Router: `/pending`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "DeletePendingInvitation",
            Router: `/pending/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "CreateTeam",
            Router: `/team/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamLeadController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamLeadController"],
        beego.ControllerComments{
            Method: "VerifiHasTeam",
            Router: `/validate`,
            AllowHTTPMethods: []string{"get"},
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
