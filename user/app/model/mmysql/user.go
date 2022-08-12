package mmysql

type UserDemo struct {
	ID        uint   `gorm:"primarykey"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (UserDemo) TableName() string {
	return "zta_user_demo"
}
