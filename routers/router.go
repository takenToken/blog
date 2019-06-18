package routers

import (
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

//初始化
func init() {


	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/register", &admin.AccountController{}, "*:Register")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")
	//系统管理
	beego.Router("/admin/system/setting", &admin.SystemController{}, "*:Setting")
	//
	////内容管理
	//beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	//beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	//beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	//beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	//beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	//beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	//beego.Router("/admin/tag", &admin.TagController{}, "*:Index")
	//
	////说说管理
	//beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")
	//beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	//beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")
	//
	////相册管理
	//beego.Router("/admin/album/add", &admin.AlbumController{}, "*:Add")
	//beego.Router("/admin/album/list", &admin.AlbumController{}, "*:List")
	//beego.Router("/admin/album/edit", &admin.AlbumController{}, "*:Edit")
	//beego.Router("/admin/album/delete", &admin.AlbumController{}, "*:Delete")
	//
	////照片管理
	//beego.Router("/admin/photo/list", &admin.PhotoController{}, "*:List")
	//beego.Router("/admin/photo/cover", &admin.PhotoController{}, "*:Cover")
	//beego.Router("/admin/photo/delete", &admin.PhotoController{}, "*:Delete")
	//
	////用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	////友链管理
	//beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	//beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")
	//beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	//beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")
	//
	////评论管理
	//beego.Router("/admin/comments/list", &admin.CommentsController{}, "*:List")
	//beego.Router("/admin/comments/add", &admin.CommentsController{}, "*:Add")
	//beego.Router("/admin/comments/edit", &admin.CommentsController{}, "*:Edit")
	//beego.Router("/admin/comments/delete", &admin.CommentsController{}, "*:Delete")
	//
	////独立fileupload
	//beego.Router("/admin/upload", &admin.FileuploadController{}, "*:Upload")
	//beego.Router("/admin/uploadfile", &admin.FileuploadController{}, "*:UploadFile")
}
