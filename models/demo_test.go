package models

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
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
		fmt.Printf("%+v", (admin)[0])
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

func TestAddAdmin(t *testing.T) {
	beforeTest()

	m2 := make(map[string]string)
	m2["username"] = randomdata.SillyName()
	m2["password"] = randomdata.Digits(9)
	admin := &Admin{0, m2["username"], m2["password"], 0}

	err := AddAdmin(admin)
	if err == nil {
		t.Logf("ok")
	} else {
		t.Error(fmt.Sprintf("%+v", err))
	}
}

func TestDeleteAdmin(t *testing.T) {
	beforeTest()

	err := DeleteAdmin("Ribshell")
	if err == nil {
		t.Logf("ok")
	} else {
		t.Error(fmt.Sprintf("%+v", err))
	}
}

func TestEditAdmin(t *testing.T) {
	beforeTest()

	m2 := make(map[string]interface{})
	m2["username"] = "tony"
	m2["password"] = "pizza"
	err := EditAdmin(m2)
	if err == nil {
		t.Logf("ok")
	} else {
		t.Error(fmt.Sprintf("%+v", err))
	}
}

//func TestQueryAdmin(t *testing.T) {
//
//	beforeTest()
//
//	m2 := make(map[string]interface{})
//	m2["username"] = "tony"
//	m2["password"] = ""
//	m2["id"] = 1
//	m2["expireday"] = 1
//
//	_adminlist,err := QueryAdmin(m2)
//
//	if err == nil {
//		t.Logf("%v",_adminlist)
//	}else{
//		t.Errorf("%v",err)
//	}
//
//}

func TestQueryAdminbyStruct(t *testing.T) {
	beforeTest()

	_admin := Admin{0, "tony", "pizza", 0}
	_adminlist, err := QueryAdminbyStruct(&_admin)

	if err == nil {
		t.Logf("%v", _adminlist)
		t.Logf("%v", (*_adminlist)[0].ID)
	} else {
		t.Errorf("%v", err)
	}
}

// 性能测试
func BenchmarkAddAdmin(b *testing.B) {

	beforeTest()
	//doSomethingPrepare(b.N)
	b.ResetTimer()

	for i := 0; i != 1000; i++ {
		m2 := make(map[string]string)
		m2["username"] = randomdata.SillyName()
		m2["password"] = randomdata.Digits(9)
		admin := &Admin{0, m2["username"], m2["password"], 0}
		err := AddAdmin(admin)
		if err == nil {
			b.Logf("ok")
		} else {
			b.Error(fmt.Sprintf("%+v", err))
		}
	}

}
