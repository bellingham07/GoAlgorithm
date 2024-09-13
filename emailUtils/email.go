package emailUtils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func emails() {
	username := "abc@163.com"
	password := "邮箱授权密码" //
	e := email.NewEmail()
	e.From = "b <" + username + ">"
	e.To = []string{"to email"}
	e.Subject = "网红商家业务对接平台"
	e.HTML = []byte("您的验证码为:<h1>" + "123456" + "</h1>" + "五分钟内有效")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", username, password, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		fmt.Println("err:", err)
	}
}
