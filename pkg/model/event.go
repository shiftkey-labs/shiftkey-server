package model

type EventStatus string

const (
	Draft     EventStatus = "draft"
	Published             = "published"
	Finished              = "finished"
)

type Attendance struct {
	ID      string `gorm:"primary_key" json:"id"`
	EventID string `json:"eventId"`
	UserID  string `json:"userId"`
	Event   Event  `gorm:"foreignkey:EventID"`
	User    User   `gorm:"foreignkey:UserID"`
}

type Event struct {
	ID                    string      `json:"id" gorm:"primary_key"`
	Title                 string      `json:"event"`
	Image                 string      `json:"img"`
	Description           string      `json:"description"`
	Location              string      `json:"location"`
	UnidentifiedAttendees int16       `json:"attendance"`
	Status                EventStatus `json:"status" gorm:"type:EventStatus"`
	Date                  string      `json:"date"`
	Time                  string      `json:"time"`
}
