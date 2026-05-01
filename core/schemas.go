package core

type SiteDataCreate struct {
	HeaderText string `json:"headertext" binding:"required"`
	BodyText   string `json:"bodytext" binding:"required"`
	ImageURL   string `json:"imageurl" binding:"required"`
	SiteURL    string `json:"siteurl" binding:"required"`
}

type LoginSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
