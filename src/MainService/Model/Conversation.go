package Model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	RoomID   uint      `gorm:"type:int;not null;index:idx_name,unique;" json:"room_id"`
	MemberID uint      `gorm:"type:int;not null;index:idx_name,unique;" json:"member_id"`
	Messages []Message `gorm:"foreignKey:ConversationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"messages"`
}

func (c *Conversation) GetID() uint {
	return c.ID
}
