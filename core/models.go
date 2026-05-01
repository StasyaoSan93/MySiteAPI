package core

type SiteData struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	HeaderText string `gorm:"column:headertext" json:"headertext"`
	BodyText   string `gorm:"column:bodytext" json:"bodytext"`
	ImageURL   string `gorm:"column:imageurl" json:"imageurl"`
	SiteURL    string `gorm:"column:siteurl" json:"siteurl"`
}

func (SiteData) TableName() string {
	return "sitedata"
}

type InfoPageData struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	HeaderText string `gorm:"column:headertext" json:"headertext"`
	BodyText   string `gorm:"column:bodytext" json:"bodytext"`
}

func (InfoPageData) TableName() string {
	return "infopagedata"
}

type User struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"id"`
	Username       string `gorm:"column:username;uniqueIndex" json:"username"`
	HashedPassword string `gorm:"column:hashed_password" json:"-"`
}

func (User) TableName() string {
	return "users"
}
