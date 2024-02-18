package model

type Device struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"-"`
	Hostname  string `json:"hostname"`
	Platform  string `json:"platform"`
	Version   string `json:"version"`
}
