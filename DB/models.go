package DB

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:64;not null"`
	Password string `gorm:"size:255;not null" json:"password,omitempty"`
}

// 一对一附属信息（UserID 既是主键又是外键）
type UserInfo struct {
	UserID    uint   `gorm:"primaryKey"` // 关联到 User.ID
	Addr      string `gorm:"size:255"`
	Email     string `gorm:"size:128"`
	Age       int
	Gender    string `gorm:"size:10"`
	Signature string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}

type Group struct {
	gorm.Model
	GroupName string `gorm:"size:100;not null"`
	OwnerID   uint   `gorm:"not null;comment:群主ID"`
	OwnerName string `gorm:"size:100;not null"`

	Owner *User `gorm:"foreignKey:OwnerID;references:ID"`
}

// 中间表
// 好友关系表（自引用多对多）
type UserUser struct {
	ActiveID  uint `gorm:"primaryKey;comment:主动添加方"`
	PassiveID uint `gorm:"primaryKey;comment:被动添加方"`
	Break     bool `gorm:"default:false;comment:是否拉黑/解除"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 关联字段
	Active  *User `gorm:"foreignKey:ActiveID;references:ID"`
	Passive *User `gorm:"foreignKey:PassiveID;references:ID"`
}

// 群成员表
type GroupUser struct {
	GroupID   uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	Break     bool `gorm:"default:false;comment:是否禁言/踢出"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 关联字段
	Group *Group `gorm:"foreignKey:GroupID;references:ID"`
	User  *User  `gorm:"foreignKey:UserID;references:ID"`
}
type GroupMessage struct {
	ID      uint `gorm:"primaryKey"`
	ToID    uint
	FromID  uint
	Content []byte

	To   *Group `gorm:"foreignKey:ToID;references:ID"`
	From *User  `gorm:"foreignKey:FromID;references:ID"`
}
type PrivateMessage struct {
	ID      uint `gorm:"primaryKey"`
	ToID    uint
	FromID  uint
	Content []byte

	To   *User `gorm:"foreignKey:ToID;references:ID"`
	From *User `gorm:"foreignKey:FromID;references:ID"`
}
