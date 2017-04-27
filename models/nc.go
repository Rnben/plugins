// GOLANG
//***********************************************
//
//      Filename: nc.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-04-25 18:11:57
// Last Modified: 2017-04-25 18:32:02
//***********************************************

package models

type NcInfo struct {
	SrcIP   string `json:"srcip,omitempty"`
	DesIP   string `json:"desip,omitempty"`
	DesPort string `json:"desport,omitempty"`
}
