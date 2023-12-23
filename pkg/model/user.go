package model

type Role string

const (
	Student   Role = "student"
	Volunteer      = "volunteer"
	Admin          = "admin"
	God            = "god"
)

type User struct {
	ID              string   `json:"id" gorm:"primary_key"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Role            Role     `json:"role" gorm:"type:Role"`
	Interests       []string `json:"interests" gorm:"type:text[]"`
	Major           string   `json:"major"`
	School          string   `json:"school"`
	CountryOfOrigin string   `json:"countryOfOrigin"`
	YearOfStudy     int16    `json:"yearOfStudy"`
	Age             int      `json:"age"`
}
