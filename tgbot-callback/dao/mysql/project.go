package mysql

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"tgbot/models"

	_ "github.com/go-sql-driver/mysql"
)

type calldb struct {
}

var Calldb = new(calldb)

// insert project link tg user
func (d *calldb) Insert() {
	projectStaff := []string{"Mark", "nike"}
	projectStaffJSON, _ := json.Marshal(projectStaff)

	_, err := db.Exec("INSERT INTO project_groups(job_name, status, project_staff,remark,cascades) VALUES (?, ?, ?, ?,?)",
		"测试", 1, string(projectStaffJSON), "is test image", "dev")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Inserted Project Alpha")
}

func (d *calldb) SelectProject(project_mark, tg_user, env string) (err error, result bool) {
	// 拆入数据
	// projectStaff := []string{"mark"}
	// projectStaffJSON, _ := json.Marshal(projectStaff)

	// _, err = db.Exec("INSERT INTO project_groups(job_name, status, project_staff) VALUES (?, ?, ?)",
	// 	"edwin", 1, string(projectStaffJSON))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("Inserted Project Alpha")

	// 查询数据
	var projects []models.ProjectGroups
	err = db.Select(&projects, "SELECT * FROM project_groups WHERE job_name = ? AND status = ? AND cascades = ?", project_mark, 1, env)
	if err != nil {
		return err, false
	}

	for _, project := range projects {
		fmt.Println(project.Remark, project.Cascades)
		// 去除双引号和方括号
		str := project.ProjectStaff.String
		str = strings.ReplaceAll(str, `"`, "")
		str = strings.ReplaceAll(str, "[", "")
		str = strings.ReplaceAll(str, "]", "")

		// 将字符串按逗号分割成切片
		slice := strings.Split(str, ", ")

		// 要检查的字符串
		// search := "Charlie"
		// 检查切片中是否存在特定的字符串
		if ContainsString(slice, tg_user) {
			return nil, true
		} else {
			return nil, false
		}
	}
	return err, false
}

// ContainsString 函数检查切片中是否包含特定的字符串
func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
