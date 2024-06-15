package controller

import (
	"fmt"
	"log"

	"slow/models"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/region"
	"gopkg.in/ini.v1"
)

func StartGet() (aks, sks, regions string, err error) {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
		return "", "", "", err
	}
	aks = cfg.Section("").Key("AccessKey").String()
	sks = cfg.Section("").Key("SecretIdAccess").String()
	regions = cfg.Section("").Key("regions").String()
	return aks, sks, regions, nil
}
func Getslow(aks, sks, regions, instances string) (sqlTime, database, user, sql, sqlStartTime, sqlType string) {
	auth := basic.NewCredentialsBuilder().
		WithAk(aks).
		WithSk(sks).
		Build()

	client := rds.NewRdsClient(
		rds.RdsClientBuilder().
			// region
			WithRegion(region.ValueOf(regions)).
			WithCredential(auth).
			Build())

	request := &model.ListSlowlogForLtsRequest{}
	databaseSlowlogForLtsRequest := "cmdb"

	// 实例ID
	request.InstanceId = instances
	xLanguageRequest := model.GetListSlowlogForLtsRequestXLanguageEnum().ZH_CN
	request.XLanguage = &xLanguageRequest
	// 单页查询慢sql条数
	limitSlowLogForLtsRequest := int32(1)

	queryStartTime, queryEndTime := models.NowStr()

	request.Body = &model.SlowlogForLtsRequest{
		Database:  &databaseSlowlogForLtsRequest,
		Limit:     &limitSlowLogForLtsRequest,
		EndTime:   queryEndTime,
		StartTime: queryStartTime,
	}

	response, err := client.ListSlowlogForLts(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 返回一个慢查询数组

	var sqlTime1, idatabase, iuser, isql, isqlStartTime, isqlType string
	slowLogList := response.SlowLogList
	for _, i2 := range *slowLogList {
		// sql 执行时间
		sqlTime1 = *i2.Time
		// 库名
		idatabase = *i2.Database
		// user
		iuser = *i2.Users
		// 详细sql语句
		isql = *i2.QuerySample
		// sql执行起始时间
		isqlStartTime = *i2.StartTime

		// sql 类型
		isqlType = *i2.Type
	}
	return sqlTime1, idatabase, iuser, isql, models.Add8(isqlStartTime), isqlType
}
