package utild

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(fromEmailAccount, fromEmailPass, toEmail, title, content string) (bool, error) {
	// 创建邮件对象
	m := gomail.NewMessage()
	// 发件人
	m.SetHeader("From", fromEmailAccount)
	// 收件人
	m.SetHeader("To", toEmail)
	// 邮件标题
	m.SetHeader("Subject", title)
	// 邮件正文
	m.SetBody("text/html", content)

	// 126邮箱SMTP服务器配置
	d := gomail.NewDialer("smtp.126.com", 465, fromEmailAccount, fromEmailPass)

	// 使用SSL连接
	d.SSL = true

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
