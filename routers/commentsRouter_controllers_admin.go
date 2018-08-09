package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "AddPost",
			Router: `/addPost`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "DeletePost",
			Router: `/deletePost/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "GetPostById",
			Router: `/getPostById/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "GetPostList",
			Router: `/getPostList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "GetPostTotal",
			Router: `/getPostTotal`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "GetPostsByCatId",
			Router: `/getPostsByCatId/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "GetPostsByTagId",
			Router: `/getPostsByTagId/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "OfflinePost",
			Router: `/offlinePost/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "PublishPost",
			Router: `/publishPost/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"] = append(beego.GlobalControllerRouter["vue-blog-api/controllers/admin:PostController"],
		beego.ControllerComments{
			Method: "UpdatePost",
			Router: `/updatePost/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
