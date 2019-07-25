package models

/*
数据库操作
*/

//type Admin struct {
//	ID        int    	`gorm:"column:ID;PRIMARY_KEY" json:"id" `
//	Username  string 	`gorm:"column:Username" json:"username" `
//	Password  string 	`gorm:"column:Password" json:"password"`
//	ExpireDay int   	`gorm:"column:Expireday" json: "expireday"`
//}

// 注意注释写法
type Admin struct {
	ID        int    `gorm:"column:ID;PRIMARY_KEY"`
	Username  string `gorm:"column:Username"`
	Password  string `gorm:"column:Password"`
	ExpireDay int    `gorm:"column:Expireday"`
}

// get data by id
// 主键查找
func GetAdminbyId(id int) (*Admin, error) {

	var ad Admin
	err := db.Where("ID = ?", id).First(&ad).Error
	if err != nil {
		return nil, err
	}
	return &ad, nil
}

// get data by name
// 多条记录
func GetAdminbyNm(nm string) ([]*Admin, error) {
	var ad []*Admin
	err := db.Where("username = ?", nm).Find(&ad).Error
	if err != nil {
		return nil, err
	}
	return ad, nil
}

// 多条件查询
func CheckAdmin(nm, pwd string) (bool, error) {
	var count int
	err := db.Model(&Admin{}).Where("Username = ? and password =?", nm, pwd).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// map方式查询
//func QueryAdmin(data map[string]interface{}) (*[]Admin, error) {
//
//	var _adminlist []Admin
//
//	_username := data["username"].(string)
//	_password := data["password"].(string)
//	_id := data["id"].(int64)
//	_expireday := data["expireday"].(int64)
//
//	queryadmin := Admin{
//		Username: sql.NullString{_username,!(_username=="")},
//		Password: sql.NullString{_password,!(_password=="")},
//		ID: sql.NullInt64{_id,!(_id==0)},
//		ExpireDay: sql.NullInt64{_expireday,!(_expireday==0)},
//	}
//
//	err := db.Where(&queryadmin).Find(&_adminlist).Error
//
//	if err != nil {
//		return nil,err
//	}else{
//		return &_adminlist,nil
//	}
//}

func QueryAdminbyStruct(admin *Admin) (*[]Admin, error) {

	var _adminlist []Admin
	err := db.Where(admin).Find(&_adminlist).Error

	if err != nil {
		return nil, err
	} else {
		return &_adminlist, nil
	}

}

// 新增数据
//func AddAdmin(data map[string]interface{}) error {
//
//	err := db.Create(&Admin{
//		Username: "",
//		Password: sql.NullString{data["password"].(string),true},
//	}).Error
//
//	return err
//}

// 新增数据
func AddAdmin(admin *Admin) error {
	err := db.Create(admin).Error
	return err
}

// 删除数据
func DeleteAdmin(nm string) error {
	return db.Where("username = ?", nm).Delete(Admin{}).Error
}

// 修改数据
func EditAdmin(data map[string]interface{}) error {
	err := db.Table("blog_admin").Where("username = ?", data["username"]).UpdateColumn("password", data["password"]).Error
	return err
}

//func isEmpty(nullString sql.NullString) bool {
//	if nullString.String == "" {
//		return true
//	}else{
//		return false
//	}
//}
//
//func isEmpty(nullString sql.NullString) bool {
//	if nullString.String == "" {
//		return true
//	}else{
//		return false
//	}
//}
