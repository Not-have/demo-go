package models

type User struct {
	Id       int    `gorm:"column:id" json:"id"` // 数据库列名 + JSON 字段名
	Username string `gorm:"column:username" json:"username"`
	Age      int    `gorm:"column:age" json:"age"`
}

func (User) TableName() string {
	return "t_user"
}
