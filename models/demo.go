package models

/*
数据库操作
*/

type Admin struct {
	ID       int    `gorm:"primary_key,column:ID " json:"id"`
	Username string `json:"username" gorm:"column:Username"`
	Password string `json:"password" gorm:"column:Password"`
}

// get data by id
// 主键查找
func GetAdminbyId(id int) (*Admin, error) {
	var ad *Admin
	err := db.Where("id = ?", id).First(&ad).Error
	if err != nil {
		return nil, err
	}
	return ad, nil
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

// 新增数据
func AddAdmin(data map[string]interface{}) error {
	err := db.Create(&Admin{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}).Error
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
