package models

import "database/sql"

type TgCallbackAcl struct {
	Id          int     `db:"id" json:"id"`
	BotName     string  `db:"bot_name" json:"bot_name"`
	TgUser      string  `db:"tg_user" json:"tg_user"`
	TgUserName  string  `db:"tg_username" json:"tg_username"`
	Status      string  `db:"status" json:"status"`
	CreatedTime string  `db:"create_time" json:"create_time"`
	UpdatedTime string  `db:"update_time" json:"update_time"`
	DeleteTime  *string `db:"delete_time" json:"delete_time"`
}
type ProjectGroups struct {
	Id           int            `db:"id" json:"id"`
	JobName      string         `db:"job_name" json:"job_name"`
	ProjectStaff sql.NullString `db:"project_staff" json:"project_staff"`
	Status       int            `db:"status" json:"status"`
	Remark       string         `db:"remark" json:"remark"`
	Cascades     string         `db:"cascades" json:"cascades"`
	CreatedTime  string         `db:"create_time" json:"create_time"`
	UpdatedTime  string         `db:"update_time" json:"update_time"`
	DeleteTime   *string        `db:"delete_time" json:"delete_time"`
}
