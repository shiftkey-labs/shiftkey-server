package model

type EventHost struct {
	ID      string `json:"id" gorm:"primary_key"`
	EventID string `json:"eventId"`
	HostID  string `json:"hostId"`
	Event   Event  `gorm:"foreignkey: EventID"`
	Host    Host   `gorm:"foreignkey: HostID"`
}

type Host struct {
	ID          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Headshot    string `json:"headshot"`
	Description string `json:"description"`
}
