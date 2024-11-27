package model

type BaseModel struct {
	Id        int64 `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt int64 `json:"created_at" gorm:"created_at;autoCreateTime:milli;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at" gorm:"updated_at;autoUpdateTime:milli;comment:更新时间"`
}
