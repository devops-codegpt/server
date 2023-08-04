package models

// Conversation saves user conversations
type Conversation struct {
	UUIDModel
	Username string     `gorm:"comment:'user name'" json:"username"`
	Title    string     `gorm:"comment:'conversation title'" json:"title"`
	State    bool       `gorm:"comment:'state'" json:"state"`
	Messages []*Message `gorm:"foreignkey:ConversationId" json:"messages"`
}
