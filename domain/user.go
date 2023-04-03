package domain

type User struct {
	ID                   string `json:"id" gorm:"type:uuid;primary_key"`
	Name                 string `json:"name" gorm:"type:varchar(255)"`
	Description          string `json:"description" gorm:"type:varchar(255)"`
	ParentServiceId      int    `json:"parentServiceId" gorm:"type:int"`
	ServiceForTicketType int    `json:"serviceForTicketType" gorm:"type:int"`
	IsVisible            int    `json:"isVisible" gorm:"type:int"`
	AllowSelection       int    `json:"allowSelection" gorm:"type:int"`
	IsActive             bool   `json:"isActive" gorm:"type:bool"`
	AutomationMacro      string `json:"automationMacro" gorm:"type:varchar(255)"`
	DefaultCategory      string `json:"defaultCategory" gorm:"type:varchar(255)"`
	DefaultUrgency       string `json:"defaultUrgency" gorm:"type:varchar(255)"`
	AllowAllCategories   bool   `json:"allowAllCategories" gorm:"type:bool"`
}

func NewUser() *User {
	return &User{}
}
