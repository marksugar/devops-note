package active

import (
	"casbin/models"
	"fmt"

	"github.com/casbin/casbin/v2"
	casbin "github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

func InsertData(db *gorm.DB, e *casbin.Enforcer) {
	// 创建角色
	role := models.Role{Name: "admin", Code: "admin"}
	db.Create(&role)

	// 创建操作
	op1 := models.Operation{Name: "Read Home", Path: "/home", Method: "GET", RoleID: role.ID}
	op2 := models.Operation{Name: "Write Data", Path: "/data", Method: "POST", RoleID: role.ID}
	db.Create(&op1)
	db.Create(&op2)

	// 创建用户并关联角色
	user := models.User{Username: "alice", Password: "password"}
	user.Roles = append(user.Roles, role)
	db.Create(&user)
	// 添加 Casbin 策略
	e.AddPolicy("admin", "/home", "GET")
	e.AddPolicy("admin", "/data", "POST")
	p := [][]string{[]string{
		"alice", "admin",
	}}
	e.AddGroupingPolicies(p)
	e.AddGroupingPolicy(p)
}

// 批量和单挑添加
func AddPolicies(e *casbin.Enforcer) {
	e.AddPolicy("user2", "/home", "GET") // 单条

	data := [][]string{
		{"user1", "/home", "GET"}, {"user1", "/data", "POST"},
	}
	// ok, err := e.AddNamedPoliciesEx("p", data)
	ok, err := e.AddPoliciesEx(data)
	if err != nil {
		fmt.Println(err)
	}
	if !ok {
		fmt.Println("添加失败")
	}
	fmt.Println("添加完成")
}
func del() {
	// 删除所有
	// e.DeleteRole("admin")

	// 删掉admin包含的/data的行
	// e.DeletePermissionForUser("admin", "/data")

	// 删掉alice1组包含admin的行
	// e.DeleteRoleForUser("alice1", "admin")

}
func AddGrouo(e *casbin.Enforcer) {
	// 添加admin对应的path
	// e.AddPolicy("admin", "/home", "GET")
	// e.AddPolicy("admin", "/data", "POST")

	// // 添加alice组对应的admin
	// e.AddGroupingPolicy("alice", "admin")

	// 添加前查询是否存在
	has, err := e.HasGroupingPolicy("alice", "user1")
	if err != nil {
		fmt.Println(err)
	}
	if !has {
		fmt.Println("不存在")
	}
	e.AddGroupingPolicy("alice", "user1")
}
func (CasbbinApp) UpdatePolicy(e *casbin.Enforcer, oldSub, oldObj, oldAct, newSub, newObj, newAct string) error {
	_, err := e.UpdatePolicy([]string{oldSub, oldObj, oldAct}, []string{newSub, newObj, newAct})
	return err
}
func (CasbbinApp) UpdateGroupingPolicy(e *casbin.Enforcer, oldr, oldg, newr, newg string) error {
	_, err := e.UpdateGroupingPolicy([]string{oldr, oldg}, []string{newr, newg})
	return err
}
func (CasbbinApp) RemovePolicy(e *casbin.Enforcer, sub, obj, act string) error {
	_, err := e.RemovePolicy(sub, obj, act)
	return err
}
func (e CasbbinApp) RemoveGroupingPolicy(oldr, oldg string) error {
	_, err := e.enforcer.RemoveGroupingPolicy(oldr, oldg)
	return err
}

type CasbbinApp struct {
	enforcer *casbin.Enforcer
}
