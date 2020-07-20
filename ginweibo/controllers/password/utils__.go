package password

import (
	"ginweibo/pkg/helpers"
	passwordResetModel "ginweibo/models/password_reset"
	"ginweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func sendResetEmail(pwd *passwordResetModel.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	resetPasswordURL := named.G("password.reset", "token", pwd.Token)
	return helpers.SendMail([]string{pwd.Email}, subject, tpl, gin.H{"resetPasswordURL": resetPasswordURL})
}
