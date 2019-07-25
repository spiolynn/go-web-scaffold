package service

import (
	"go-web-scaffold/models"
	"go-web-scaffold/pkgs/util"
)

/*
@feature: 查询admin信息列表
@in: adminmap
@out: json string
@steps:
	1 adminmap -> struct
	2 call model.QueryList
	3 []struct -> json
*/

func QueryAdminList(adminmap map[string]interface{}) (string, error) {

	// 返回
	data := ""

	// adminmap -> struct
	var a_admin models.Admin
	util.Map_to_Struct(adminmap, &a_admin)

	// call model
	_adminlist, err := models.QueryAdminbyStruct(&a_admin)

	if err != nil {
		return data, err
	} else {
		// []struct to json
		json_str, _ := util.Struct_to_Json(_adminlist)
		return json_str, nil
	}

}
