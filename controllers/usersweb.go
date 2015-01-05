package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/astaxie/beego"
)

type UsersWebController struct {
	beego.Controller
}

func (this *UsersWebController) Prepare() {
	beego.Debug(fmt.Sprintf("[%s] %s | %s", this.Ctx.Input.Host(), this.Ctx.Input.Request.Method, this.Ctx.Input.Request.RequestURI))
	beego.Debug("[Header] ")
	beego.Debug(this.Ctx.Request.Header)
}

func (this *UsersWebController) PostGravatar() {
	var result Result
	file, fileHeader, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		result = Result{Success: false}
		this.Data["json"] = &result
		this.ServeJson()
		return
	}
	f, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//处理文件错误
		fmt.Println("OpenFile Error")
	}
	defer f.Close()
	io.Copy(f, file)
	this.Ctx.Output.Context.Output.SetStatus(http.StatusOK)
	this.Ctx.Output.Context.ResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")

	result = Result{Success: true}
	this.Data["json"] = &result
	this.ServeJson()
}
