package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slow/config"
	"slow/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"gopkg.in/ini.v1"
)

func IsId(instances string) (ok bool) {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	fmt.Println(instances, cfg.Section("").Key("instances").String())
	return instances == cfg.Section("").Key("instances").String()
}
func PostJson(c *gin.Context) {
	// 处理逻辑
	var message *config.HuaweicloudSlowPostMessage
	c.BindJSON(&message)
	messageMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(message.Message), &messageMap)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return
	}
	rID := gjson.GetBytes([]byte(message.Message), "template_variable.ResourceId").Str
	sms_content := gjson.GetBytes([]byte(message.Message), "sms_content").Str
	metric_name := gjson.GetBytes([]byte(message.Message), "metric_name").Str
	alarm_id := gjson.GetBytes([]byte(message.Message), "alarm_id").Str

	// HuaweiMonLists := []string{"rds001_cpu_util", "rds039_disk_util", "rds002_mem_util", "rds074_slow_queries"}
	logrus.WithFields(logrus.Fields{
		"input message: ":     metric_name,
		"input sms_content: ": sms_content,
		"input alarm_id: ":    alarm_id,
	}).Debug("triggerMessage collector Request.Body")

	// rds074_slow_queries是华为云官方的结构，详情见官方slow文档
	switch metric_name {
	case "rds074_slow_queries":
		SlowInfo(rID, sms_content)
		return
	default:
		models.SedMessageDingtalkInfo(sms_content, alarm_id)
		logrus.WithFields(logrus.Fields{
			"illegal input message: ": nil,
		}).Debug("triggerMessage collector Request.Body")
		return
	}
}

func SlowInfo(rID, sms_content string) {
	aks, sks, regions, err := StartGet()
	if err != nil {
		panic(err)
	}
	if IsId(rID) {
		sqlTime, database, user, sql, sqlStartTime, sqlType := Getslow(aks, sks, regions, rID)
		if sqlTime == "" {
			logrus.WithFields(logrus.Fields{
				"triggerMessage collector Getslow reust zero: ":  sqlTime,
				"is not ResourceId or Other invalid information": rID,
			}).Debug("triggerMessage collector")
			return
		}
		models.SedMessageDingtalk(sms_content, sqlTime, database, user, sql, sqlStartTime, sqlType)
		logrus.WithFields(logrus.Fields{
			"rId":          rID,
			"sms_content":  sms_content,
			"sqlTime":      sqlTime,
			"database":     database,
			"user":         user,
			"sql":          sql,
			"sqlStartTime": sqlStartTime,
			"sqlType":      sqlType,
		}).Debug("triggerMessage collector")
	}
}
