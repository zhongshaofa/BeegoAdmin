package models

import "time"

type SystemAdmin struct {
	Id       int
	Username string
	Password string
	Phone    string
	Remark   string
	Status   int
	CreateAt time.Time `orm:"auto_now;type(datetime)"`
}

func (m *SystemAdmin) TableName() string {
	return TableName("system_admin")
}
