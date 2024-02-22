package dao

func QueryGroupPermissionDataInPage(pageSize int, offset int) []GroupPermission {
	var users []GroupPermission
	DBClient.Limit(pageSize).Offset(offset).Find(&users)
	return users
}
