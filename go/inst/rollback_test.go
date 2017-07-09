package inst

import (
	test "github.com/openark/golib/tests"
	"testing"
)

func TestAddRollbackAction(t *testing.T) {
	AddRollbackAction("start_slave", "mydb1:3306")
	test.S(t).ExpectEquals(rollbackActions[0], rollbackAction{"start_slave", "mydb1:3306"})
}

func TestRollback(t *testing.T) {
	rollbackActions = []rollbackAction{}
	AddRollbackAction("xxxxxx", "mydb1:3306")
	err := Rollback()
	test.S(t).ExpectNotNil(err)
}
