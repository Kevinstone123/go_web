package dbops

import "testing"

func TestMain(m *testing.M){
	m.Run()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("testGetUser", testGetUser)
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("1101")
	if pwd != "市辖区" ||  err != nil {
		t.Errorf("Error of GetUser")
	}
}