package dao

func GetFollowCount(userId int64) int64 {
	var followCount int64
	db.Select("Count(*)").Where("userid = ?", userId).Table("follow_info").Find(&followCount)
	return followCount
}
