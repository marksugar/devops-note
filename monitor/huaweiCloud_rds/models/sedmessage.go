package models

import (
	"fmt"
	"log"

	"github.com/biello/dingtalk-webhook-client/client"
	"gopkg.in/ini.v1"
)

func ConfigINI() (accessToken, secret, dingtalklist1 string) {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	accessToken = cfg.Section("").Key("DingtalkWebHookKey").String()
	secret = cfg.Section("").Key("DingtalkWebHooKSEC").String()
	dingtalklist1 = cfg.Section("").Key("dingtalklist").String()
	return accessToken, secret, dingtalklist1
}
func SedMessageDingtalk(sms_content, sqlTime, database, user, sql, sqlStartTime, sqlType string) {
	accessToken, secret, dingtalklist1 := ConfigINI()
	cli := client.DefaultDingTalkClient(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken), secret)
	markDownReq := client.OapiRobotSendRequest{
		MsgType: "markdown",
		Markdown: client.Markdown{
			Title: "## MySQL Slow Info",
			Text: "#### **MySQL Slow Info**\n" +
				fmt.Sprintf("> ### 执行时间(s): %s\n", sqlTime) +
				fmt.Sprintf("> ### 执行库: %s\n", database) +
				fmt.Sprintf("> ### 帐号: %s\n", user) +
				fmt.Sprintf("> ### 发生时间: %s\n", sqlStartTime) +
				fmt.Sprintf("> ### 语句类型: %s\n", sqlType) +
				fmt.Sprintf("> #### 关联信息: %s\n", sms_content) +
				fmt.Sprintf("> ### SQL: \n%s", sql),
		},
		At: client.At{
			AtMobiles: []string{fmt.Sprintf("%s", dingtalklist1)},
			IsAtAll:   true,
		},
	}
	_, err := cli.Execute(markDownReq)
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}

}

func SedMessageDingtalkInfo(sms_content, alarm_id string) {
	accessToken, secret, dingtalklist1 := ConfigINI()
	cli := client.DefaultDingTalkClient(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken), secret)
	markDownReq := client.OapiRobotSendRequest{
		MsgType: "markdown",
		Markdown: client.Markdown{
			Title: "## MySQL Slow Info",
			Text: "#### **Alert Info**\n" +
				fmt.Sprintf("> #### info: %s\n", sms_content) +
				fmt.Sprintf("> ### 关联ID: %s", alarm_id),
		},
		At: client.At{
			AtMobiles: []string{fmt.Sprintf("%s", dingtalklist1)},
			IsAtAll:   true,
		},
	}
	_, err := cli.Execute(markDownReq)
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}
}
