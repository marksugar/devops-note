package db

import (
	"log/slog"
	"orm/models"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// create DATABASE orm_test charset utf8mb4;
var DB *gorm.DB

// https://gorm.io/zh_CN/docs/gorm_config.html#%E5%91%BD%E5%90%8D%E7%AD%96%E7%95%A5
func Initdb() *gorm.DB {

	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN:                       "test:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DSN:                       "root:1234@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{ // 是否附属表名.命名策略
			// TablePrefix:   "test_",                           // 表前缀，如果是 `User` 那么就应该是 `t_users`
			SingularTable: true,                              // 使用单数表名，启用此选项后，`User` 的表将是 `user`
			NoLowerCase:   false,                             // 跳过名称的蛇形大小写 ,不跳过则自动转换
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在将其转换为数据库名称之前，使用名称替换器更改结构体/字段名称
		},
		// https://gorm.io/zh_CN/docs/gorm_config.html#DisableForeignKeyConstraintWhenMigrating
		DisableForeignKeyConstraintWhenMigrating: true, // 外键约束是否禁用，禁用后不会建立物理外键，物理外键会导致查询缓慢，以逻辑外键为主
	})

	if err != nil {
		panic("failed to connect database")
	}

	mysql, _ := db.DB()
	//设置连接池
	// https://gorm.io/zh_CN/docs/dbresolver.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
	// https://gorm.io/zh_CN/docs/generic_interface.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
	mysql.SetMaxIdleConns(100) //连接池最大空闲俩连接数
	mysql.SetMaxOpenConns(200) // 数据库最大连接数
	mysql.SetConnMaxIdleTime(time.Hour)
	mysql.SetConnMaxLifetime(24 * time.Hour)

	//GORM.AutoMigrate(po.DeployLog{})
	slog.Info("连接数据库成功!")

	DB = db
	// register()
	return db
}
func Close(db *gorm.DB) error {
	slog.Info("关闭数据库连接")
	mysql, err := db.DB()
	if err != nil {
		return err
	}
	return mysql.Close()
}
func register() {
	// 注册表
	// https://gorm.io/zh_CN/docs/migration.html

	// 1.自动注册表
	// DB.AutoMigrate(&models.User{})

	// 2.手动注册
	m := DB.Migrator()

	// 如果表不存在就创建表
	if !m.HasTable(&models.User{}) { // 查询是否存在
		slog.Info("创建User表")
		m.CreateTable(&models.User{}) // 创建
	}
	if !m.HasTable(&models.UserTwo{}) { // 查询是否存在
		slog.Info("创建UserTwo表")
		m.CreateTable(&models.UserTwo{}) // 创建
	}
	if !m.HasTable(&models.UserInfo{}) { // 查询是否存在
		slog.Info("创建UserInfo表")
		m.CreateTable(&models.UserInfo{}) // 创建
	}

	// // 如果表存在就修改表名
	// if m.HasTable(&models.User{}) { // 查询是否存在
	// 	fmt.Println("修改表名")
	// 	m.RenameTable(&models.User{}, &models.UserTwo{})
	// }
	// 删除表
	// if err = m.DropTable(&User{}); err != nil {
	// 	fmt.Println("删除失败")
	// }
}
