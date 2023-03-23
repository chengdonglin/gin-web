package tabel

type User struct {
	Id          int64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键'"`
	Username    string `gorm:"column:username;index:,unique;type:varchar(32);not null;default:'';comment:'用户名'"`
	Password    string `gorm:"column:password;type:varchar(64);not null;default:'';comment:'密码'"`
	UpdatedTime int64  `gorm:"column:create_time;autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	CreatedTime int64  `gorm:"column:update_time;autoCreateTime"`       // 使用时间戳秒数填充创建时间
}

func (*User) TableName() string {
	return "tb_user"
}
