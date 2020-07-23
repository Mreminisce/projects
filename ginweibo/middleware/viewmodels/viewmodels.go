package viewmodels

import (
	statusModel "ginweibo/models/status"
	userModel "ginweibo/models/user"
	"ginweibo/utils/time"
)

// 用户
type UserViewModel struct {
	ID      int
	Name    string
	Email   string
	Avatar  string
	IsAdmin bool
}

// 微博
type StatusViewModel struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt string
}

// 用户数据展示
func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
	return &UserViewModel{
		ID:      int(u.ID),
		Name:    u.Name,
		Email:   u.Email,
		Avatar:  u.Gravatar(),
		IsAdmin: u.IsAdminRole(),
	}
}

// 微博数据展示
func NewStatusViewModelSerializer(s *statusModel.Status) *StatusViewModel {
	return &StatusViewModel{
		ID:        int(s.ID),
		Content:   s.Content,
		UserID:    int(s.UserID),
		CreatedAt: time.SinceForHuman(s.CreatedAt),
	}
}
