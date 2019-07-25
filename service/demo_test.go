package service

import (
	"go-web-scaffold/models"
	"testing"
)

func before_test() {
	models.OpenDB()
}

func after_test() {
	models.CloseDB()
}

func TestQueryAdminList(t *testing.T) {

	before_test()
	defer after_test()

	m2 := make(map[string]interface{})
	m2["username"] = "tony"
	m2["password"] = ""

	_adminmap, err := QueryAdminList(m2)

	if err == nil {
		t.Logf("%v", _adminmap)
	} else {
		t.Errorf("%v", err)
	}
}
