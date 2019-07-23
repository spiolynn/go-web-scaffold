package models

import (
	"fmt"
	"testing"
)

/*
TDD
*/

func beforeTest() {
	_init()
}

func TestGetAdminbyId(t *testing.T) {

	beforeTest()

	admin, err := GetAdminbyId(1)
	if err == nil {
		fmt.Printf("%+v", *admin)
	} else {
		fmt.Printf("%v", err)
	}

}

func TestGetAdminbyNm(t *testing.T) {
	beforeTest()

	admin, err := GetAdminbyNm("tony")
	if err == nil {
		fmt.Printf("%+v", admin)
	} else {
		fmt.Printf("%v", err)
	}

}

func TestCheckAdmin(t *testing.T) {
	beforeTest()

	isexist, err := CheckAdmin("tony", "teacher")
	if err != nil {
		t.Error("error")
	} else {
		if isexist == true {
			t.Log("ok")
		} else {
			t.Log("fail")
		}
	}
}
