package main

import (
	"casbin/models"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	casbinmodels "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CasbinAuthRequest struct {
	Username string `json:"username"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

// CasbinMiddleware is a middleware for Casbin authorization
func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	if err := e.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}
	return func(c *gin.Context) {
		var request CasbinAuthRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(c.Request.URL.Path, c.Request.Method, request.Username)

		obj := c.Request.URL.Path // 请求路径
		act := c.Request.Method   // 请求方法

		ok, err := e.Enforce(request.Username, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Authorization error"})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		c.Next()
	}

}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN:                       "test:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DSN:               "root:1234@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256, // string 类型字段的默认长度
	}))
	if err != nil {
		panic("failed to connect database")
	}
	mysql, _ := db.DB()
	defer mysql.Close()
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Operation{})
	// 定义结构体
	// adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &models.CasbinPolicy{})
	// e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// 初始化 Casbin 适配器和策略
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Failed to create Gorm adapter: %v", err)
	}
	// e, err := casbin.NewEnforcer("models.conf", a)
	m, err := casbinmodels.NewModelFromString(`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	// 创建 Casbin Enforcer
	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}
	if err = e.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}

	// 插入数据
	// active.InsertData(db, e)

	// // 修改alice组到alice1
	// ress, err := e.UpdateGroupingPolicy([]string{"alice"}, []string{"alice1"})
	// if err != nil {
	// 	log.Fatalf("UpdateGroupingPolicy: %v", err)
	// }
	// if !ress {
	// 	// 如果修改的值不存在，则失败
	// 	fmt.Println("修改失败")
	// }

	// active.AddPolicies(e) // 添加用户和path，method
	// active.AddGrouo(e)    // 添加组和用户关系
	// active.UpdatePolicy(e, "user2", "/home", "GET", "user1", "/test", "GET")  // 更新用户的path和method
	// active.UpdateGroupingPolicy(e, "alice", "user1", "alice", "user2")  // 更新校色和group
	// active.RemovePolicy(e, "user1", "/test", "GET") // 删除用户path,method
	// active.RemoveGroupingPolicy(e, "alice", "user1")
	e.RemoveGroupingPolicy("alice", "user2")

	// 初始化 Gin 路由
	r := gin.Default()

	// 添加中间件到路由
	r.Use(CasbinMiddleware(e))

	// 定义路由
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Home"})
	})

	r.POST("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Data"})
	})

	r.Run(":8080")
}
