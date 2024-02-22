package dao

type GroupPermission struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
	Path     string `json:"path"`
	Api      string `json:"api"`
	PathType string `json:"pathType"`
}

func (gp *GroupPermission) TableName() string {
	return "group_permission"
}
