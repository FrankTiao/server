// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameWebSession = "chii_os_web_sessions"

// WebSession mapped from table <chii_os_web_sessions>
type WebSession struct {
	Key       string `gorm:"column:key;type:char(64);primaryKey" json:"key"`               // session key
	UserID    uint32 `gorm:"column:user_id;type:int(10) unsigned;not null" json:"user_id"` // uint32 user id
	Value     []byte `gorm:"column:value;type:mediumblob;not null" json:"value"`           // json encoded session data
	CreatedAt int64  `gorm:"column:created_at;type:bigint(20);not null" json:"created_at"` // int64 unix timestamp, when session is created
	ExpiredAt int64  `gorm:"column:expired_at;type:bigint(20);not null" json:"expired_at"` // int64 unix timestamp, when session is expired
}

// TableName WebSession's table name
func (*WebSession) TableName() string {
	return TableNameWebSession
}
