package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getAll`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/getOne/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
