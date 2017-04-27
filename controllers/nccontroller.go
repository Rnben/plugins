package controllers

// GOLANG
//***********************************************
//
//      Filename: nccontroller.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-04-25 18:15:32
// Last Modified: 2017-04-27 16:57:40
//***********************************************

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xwisen/plugins/models"
	"os/exec"
	"strings"
)

type NcController struct {
	beego.Controller
}

func (ncc *NcController) TelnetPost() {
	logs.Info("Start nc .......................")
	logs.Info("request body is : %s", ncc.Ctx.Input.URI())
	if len(ncc.Ctx.Input.RequestBody) == 0 {
		logs.Error("request body is null !")
		return
	}
	nc := models.NcInfo{}
	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, &nc)
	if err != nil {
		logs.Error("Parse request body to ncinfo err : %s", err)
		return
	}
	logs.Info("nc is: %+v", nc)
	//ssh 10.78.221.181 "nc 10.78.182.10 2181 -i 1 -w 1 -v"
	ncCMD := exec.Command("ssh", nc.SrcIP, "nc", nc.DesIP, nc.DesPort, "-i", "1", "-w", "1", "-v")
	logs.Info("ssh %s nc %s %s -i 1 -w 1 -v", nc.SrcIP, nc.DesIP, nc.DesPort)
	s, err := ncCMD.CombinedOutput()
	if err != nil {
		//result := strings.SplitN(string(s), "\n", 2)
		logs.Error("exec command err : %s;result is:%s", err, strings.Split(string(s), "\n"))
		if strings.Contains(string(s), "Idle timeout expired (1000 ms)") {
			ncc.Data["json"] = "{\"info\":\"连接成功\"}"
			ncc.ServeJSON()
			return
		} else {
			ncc.Data["json"] = "{\"info\":\"连接失败\"}"
			ncc.ServeJSON()
			return
		}
	}
	//result := strings.SplitN(string(s), "\n", 2)
	logs.Info("exec command result is : %s", strings.Split(string(s), "\n"))
	if strings.Contains(string(s), "Idle timeout expired (1000 ms)") {
		ncc.Data["json"] = "{\"info\":\"连接成功\"}"
		ncc.ServeJSON()
		return
	} else {
		ncc.Data["json"] = "{\"info\":\"连接失败\"}"
		ncc.ServeJSON()
		return
	}
}

func (ncc *NcController) TelnetGet() {

}

func (ncc *NcController) GetVersion() {
	ncc.Data["json"] = "{\"version\":\"v1.0\"}"
	ncc.ServeJSON()

}
