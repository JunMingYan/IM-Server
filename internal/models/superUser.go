package models

type SuperUser struct {
	SID      int64
	Account  string
	Password string
	Role     int    //角色分类：超级管理员0，具有增删改查权限；普通管理员1，只有有查的权限
	NickName string //= "wzomg-admin";
	Avatar   string //= "img/admin-avatar.gif";
}
