package status

import (
	"ginweibo/app/controllers"
	userModel "ginweibo/app/models/user"

	"github.com/gin-gonic/gin"
)

func backTo(c *gin.Context, currentUser *userModel.User) {
	back := c.DefaultPostForm("back", "")
	if back != "" {
		controllers.Redirect(c, back, true)
		return
	}
	controllers.RedirectRouter(c, "users.show", currentUser.ID)
}
