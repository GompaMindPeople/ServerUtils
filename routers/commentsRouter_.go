package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "GetGameServerConfigForJson",
            Router: `/ManUtils/getGameServerConfigForJson`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "DeleteById",
            Router: `/deleteById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "ExecuteShell",
            Router: `/executeShell`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "ListForButtonConf",
            Router: `/listForButtonConf`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
        beego.ControllerComments{
            Method: "SaveForButtonConf",
            Router: `/saveAddDialog`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})



}
