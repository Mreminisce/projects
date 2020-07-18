package viewmodels

import (
	statusModel "ginweibo/app/models/status"
	"ginweibo/pkg/time"
)

// 微博
type StatusViewModel struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt string
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
