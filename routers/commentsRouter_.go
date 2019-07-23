package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "GetGameServerConfigForJson",
			Router:           `/ManUtils/getGameServerConfigForJson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "DeleteById",
			Router:           `/deleteById`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "DeleteGroupById",
			Router:           `/deleteGroupById`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "ExecuteShell",
			Router:           `/executeShell`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "ListForButtonConf",
			Router:           `/listForButtonConf`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:MainUtilController"],
		beego.ControllerComments{
			Method:           "ListGroupALL",
			Router:           `/listGroupALL`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"],
		beego.ControllerComments{
			Method:           "GetSshConfig",
			Router:           `/getSshConfigById`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"],
		beego.ControllerComments{
			Method:           "ListSSHCombobox",
			Router:           `/listSSHCombobox`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:SSHController"],
		beego.ControllerComments{
			Method:           "SaveSSHConfig",
			Router:           `/saveSSHConfig`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "DownloadConfigAllTemplate",
			Router:           `/downloadConfigAllTemplate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("fileName"),
			),
			Filters: nil,
			Params:  nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "GenerateConfigAllTemplate",
			Router:           `/generateConfigAllTemplate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "GetTemplateById",
			Router:           `/getTemplateById`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "ListServerName",
			Router:           `/listServerName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "SaveConfigAll",
			Router:           `/saveConfigAll`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"] = append(beego.GlobalControllerRouter["ServerUtils/controllers:TemplateController"],
		beego.ControllerComments{
			Method:           "SaveServers",
			Router:           `/saveServers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
