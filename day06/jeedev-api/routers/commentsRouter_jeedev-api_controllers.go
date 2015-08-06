package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
