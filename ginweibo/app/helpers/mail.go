package helpers

import (
	"ginweibo/config"
	"ginweibo/pkg/file"
	"ginweibo/pkg/mail"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, templatePath string, tplData map[string]interface{}) error {
	filePath := config.AppConfig.ViewsPath + "/" + templatePath
	body, err := file.ReadTemplateToString(templatePath, filePath, tplData)
	if err != nil {
		return err
	}
	mail := &mail.Mail{
		Driver:   config.MailConfig.Driver,
		Host:     config.MailConfig.Host,
		Port:     config.MailConfig.Port,
		User:     config.MailConfig.User,
		Password: config.MailConfig.Password,
		FromName: config.MailConfig.FromName,
		MailTo:   mailTo,
		Subject:  subject,
		Body:     body,
	}
	return mail.Send()
}
