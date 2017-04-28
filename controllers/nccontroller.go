package controllers

// GOLANG
//***********************************************
//
//      Filename: nccontroller.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-04-25 18:15:32
// Last Modified: 2017-04-28 10:39:30
//***********************************************

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xwisen/plugins/models"
	"os/exec"
	"regexp"
	"strings"
)

const (
	ipregexp   string = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	portregexp string = `^([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-5]{2}[0-3][0-5])$`
)

type NcController struct {
	beego.Controller
}

func (ncc *NcController) TelnetPost() {
	ncc.TplName = "telnet.tpl"
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
	//匹配IP信息
	IPRegexp := regexp.MustCompile(ipregexp)
	if err != nil {
		logs.Error("匹配IP的正则表达式不正确")
		ncc.Data["json"] = "{\"info\":\"匹配IP的正则表达式不正确\"}"
		ncc.ServeJSON()
		return
	}
	matched := IPRegexp.MatchString(nc.SrcIP)
	if !matched {
		logs.Error("源地址格式不正确")
		ncc.Data["json"] = "{\"info\":\"源地址格式不正确\"}"
		ncc.ServeJSON()
		return
	}
	matched = IPRegexp.MatchString(nc.DesIP)
	if !matched {
		logs.Error("目的地址格式不正确")
		ncc.Data["json"] = "{\"info\":\"目的地址格式不正确\"}"
		ncc.ServeJSON()
		return
	}
	//匹配端口信息
	PortRegexp := regexp.MustCompile(portregexp)
	if err != nil {
		logs.Error("匹配端口的正则表达式不正确")
		ncc.Data["json"] = "{\"info\":\"匹配端口的正则表达式不正确\"}"
		ncc.ServeJSON()
		return
	}
	matched = PortRegexp.MatchString(nc.DesPort)
	if !matched {
		logs.Error("目的端口格式不正确")
		ncc.Data["json"] = "{\"info\":\"目的端口格式不正确\"}"
		ncc.ServeJSON()
		return
	}
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
	ncc.TplName = "telnet.tpl"
}

func (ncc *NcController) GetVersion() {
	ncc.Data["json"] = "{\"version\":\"v1.0\"}"
	ncc.ServeJSON()

}
