package global

import "time"

// BaseModel Base model struct, not including deleted_at field
type BaseModel struct {
	ID        uint       `gorm:"primarykey" json:"ID"` // Primary key ID
	CreatedAt time.Time  `json:"createdAt"`            // Create time
	UpdatedAt time.Time  `json:"updatedAt"`            // Update time
	DeleteAt  *time.Time `gorm:"index" json:"-"`       // Delete time, not included in JSON
}
