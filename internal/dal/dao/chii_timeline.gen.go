// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNameTimeLine = "chii_timeline"

// TimeLine mapped from table <chii_timeline>
type TimeLine struct {
	ID       model.TimeLineID `gorm:"column:tml_id;type:int(10) unsigned;primaryKey;autoIncrement:true"`
	UID      model.UserID     `gorm:"column:tml_uid;type:mediumint(8) unsigned;not null"`
	Cat      uint16           `gorm:"column:tml_cat;type:smallint(6) unsigned;not null"`
	Type     uint16           `gorm:"column:tml_type;type:smallint(6) unsigned;not null"`
	Related  string           `gorm:"column:tml_related;type:char(255);not null;default:0"`
	Memo     []byte           `gorm:"column:tml_memo;type:mediumtext;not null"`
	Img      []byte           `gorm:"column:tml_img;type:mediumtext;not null"`
	Batch    uint8            `gorm:"column:tml_batch;type:tinyint(3) unsigned;not null"`
	Source   uint8            `gorm:"column:tml_source;type:tinyint(3) unsigned;not null"`    // 更新来源
	Replies  uint32           `gorm:"column:tml_replies;type:mediumint(8) unsigned;not null"` // 回复数
	Dateline uint32           `gorm:"column:tml_dateline;type:int(10) unsigned;not null"`
}

// TableName TimeLine's table name
func (*TimeLine) TableName() string {
	return TableNameTimeLine
}
