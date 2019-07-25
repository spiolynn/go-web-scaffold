package v1

/*
接口数据
*/

type DQueryAdminList struct {
	Username string ` json:"username" valid:"Required; MaxSize(50)" `
	Password string ` json:"password" MaxSize(50)" `
	ID       int    ` json:"id" `
}
