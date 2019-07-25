package v1

/*
接口数据
*/

//type QueryAdmin struct {
//	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
//	Password string `form:"password" json:"password" xml:"password" binding:"required"`
//}

type Admins struct {
	Username string `json:"username" valid:"Required;MaxSize(50)"`
	Password string `json:"password" valid:"MaxSize(50)"`
}
