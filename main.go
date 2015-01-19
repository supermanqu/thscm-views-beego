package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"tone-world.com/controllers"
	//	"tone-world.com/models"
	_ "tone-world.com/routers"
)

var FilterUser = func(ctx *context.Context) {
	if ctx.Request.RequestURI == "/login" {
		return
	}
	sid := ctx.GetCookie("session.toneworld.com")
	if sid == "" {
		ctx.Redirect(302, "/login")
		return
	}

	sess, error := controllers.GlobalSessions.GetSessionStore(sid)
	if error != nil {

	}

	if sess != nil {
		uid := sess.Get("uid")
		if uid == nil || uid == "" {
			ctx.Redirect(302, "/login")
		}
		//user := sess.Get("user")
		//if _, ok := user.(*models.User); ok {

		//} else {
		//	ctx.Redirect(302, "/login")
		//}
	} else {
		ctx.Redirect(302, "/login")
	}
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.SessionOn = true
	//beego.SessionProvider = "memory"

	//beego.SessionProvider = "mysql"

	//beego.SessionGCMaxLifetime = 60 * 60 //60 seconds
	//beego.SessionName = "session.toneworld.com"
	beego.SessionCookieLifeTime = -1 //60 seconds
	beego.SessionAutoSetCookie = true

	beego.InsertFilter("/", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.SetStaticPath("/scripts", "views/scripts")
	beego.SetStaticPath("/styles", "views/styles")
	beego.SetStaticPath("/fonts", "views/fonts")
	beego.SetStaticPath("/images", "views/images")
	beego.SetStaticPath("/views", "views/views")

	beego.Run()
}
