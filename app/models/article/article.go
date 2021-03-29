package article

import (
	"github.com/wuyan94zl/api/app/models/admin"
	"time"
)

type Article struct {
	Id          uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"json:"id"`
	Title       string `validate:"required,min:10,max:50"search:"like"json:"title"`
	Description string `validate:"required,min:10,max:200"json:"description"`
	Content     string `validate:"required"json:"content"`
	View        uint64
	AdminId     uint64      `validate:"required,numeric"json:"admin_id"`
	Admin       admin.Admin `gorm:"-"relationship:"belongTo"json:"admin"`
	CreatedAt   time.Time   `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at"json:"updated_at"`
}
