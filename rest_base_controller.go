package beecmf

import (
	"github.com/astaxie/beego"
)

type RestBaseController struct {
	beego.Controller
}

// 成功返回, 用法
// Sucsess("成功信息")
// Success("成功信息","跳转链接")
// Success("成功信息",返回数据struct)
// Success("成功信息",返回数据struct,"跳转链接")
func (this *RestBaseController) Success(message string, data ...interface{}) {

	var (
		url string
		mData interface{}
		dataLength int
	)

	dataLength = len(data)

	if (dataLength > 1) {

		mData = data[0]

		if v, ok := data[1].(string); ok {
			url = v
		}
	}

	if (dataLength ==1) {
		if v, ok := data[0].(string); ok {
			url = v
		}else{
			mData= data[0]
		}

	}

	result := struct {
		Message string
		Url     string
		Data    interface{}
		Status  int
	}{
		Url:url,
		Data:mData,
		Message:message,
		Status:1,
	}

	this.Data["json"] = result

	if (beego.BConfig.RunMode == beego.DEV) {
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	this.ServeJSON()

}


// 成功返回, 用法
// Error("错误信息")
// Error("错误信息","跳转链接")
// Error("错误信息",返回数据struct)
// Error("错误信息",返回数据struct,"跳转链接")
func (this *RestBaseController) Error(message string, data ...interface{}) {

	var (
		url string
		mData interface{}
		dataLength int
	)

	dataLength = len(data)

	if (dataLength > 1) {

		mData = data[0]

		if v, ok := data[1].(string); ok {
			url = v
		}
	}

	if (dataLength ==1) {
		if v, ok := data[0].(string); ok {
			url = v
		}else{
			mData= data[0]
		}

	}

	result := struct {
		Message string
		Url     string
		Data    interface{}
		Status  int
	}{
		Url:url,
		Data:mData,
		Message:message,
		Status:0,
	}

	this.Data["json"] = result

	if (beego.BConfig.RunMode == beego.DEV) {
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	this.ServeJSON()

}