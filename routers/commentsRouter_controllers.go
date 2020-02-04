package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "DeleteKPI",
            Router: `/:kpiid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "GetKPIFromID",
            Router: `/:kpiid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "AssignKPI",
            Router: `/assign/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "GetKPIRange",
            Router: `/range/`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "GetMemberKPIReport",
            Router: `/report/:userid/:month`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "ScoreKPI",
            Router: `/score/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "GetAllTasks",
            Router: `/task/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "CreateTask",
            Router: `/task/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "GetAllKPITasks",
            Router: `/task/:kpiid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:KPIController"] = append(beego.GlobalControllerRouter["performance/controllers:KPIController"],
        beego.ControllerComments{
            Method: "MarkTaskComplete",
            Router: `/task/complete/:tid`,
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:ObjectiveController"] = append(beego.GlobalControllerRouter["performance/controllers:ObjectiveController"],
        beego.ControllerComments{
            Method: "CreateObjective",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:ObjectiveController"] = append(beego.GlobalControllerRouter["performance/controllers:ObjectiveController"],
        beego.ControllerComments{
            Method: "GetTeamStrategiveObjectives",
            Router: `/lead/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:ObjectiveController"] = append(beego.GlobalControllerRouter["performance/controllers:ObjectiveController"],
        beego.ControllerComments{
            Method: "GetMemberStrategiveObjectives",
            Router: `/member/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "DeleteTeam",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
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
            Method: "TakeBehaviourTest",
            Router: `/behaviour/`,
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

    beego.GlobalControllerRouter["performance/controllers:TeamController"] = append(beego.GlobalControllerRouter["performance/controllers:TeamController"],
        beego.ControllerComments{
            Method: "VerifyHasTeam",
            Router: `/verifi/:userid`,
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
