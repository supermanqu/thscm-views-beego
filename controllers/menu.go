package controllers

import (
	"github.com/astaxie/beego"
	"tone-world.com/models/mongo"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) Get() {

	conn := mongo.Conn()
	defer conn.Close()

	collection := conn.DB("thscm").C("menu")

	var result []map[string]interface{}
	iter := collection.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		return
	}

	this.Ctx.Output.Json(result, true, false)
}
