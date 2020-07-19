package factory

import (
	followerModel "ginweibo/models/follower"
	userModel "ginweibo/models/user"
)

func FollowerTableSeeder(needCleanTable bool) {
	if needCleanTable {
		DropAndCreateTable(&followerModel.Follower{})
	}
	users, err := userModel.All()
	if err != nil {
		panic("follower mock error!")
	}
	user := users[0]
	userID := user.ID
	// 获取除了 ID 为 0 的其他所有用户 ID
	followers := users[1:]
	followerIDs := make([]uint, 0)
	for _, v := range followers {
		followerIDs = append(followerIDs, v.ID)
	}
	// 关注除了 0 号用户以外的所有用户
	followerModel.DoFollow(userID, followerIDs...)
	// 除了 0 号用户以外的所有用户都关注 1 号用户
	for _, v := range followerIDs {
		followerModel.DoFollow(v, userID)
	}
}
