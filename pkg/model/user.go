package model

type Role string

const (
	Student   Role = "student"
	Volunteer      = "volunteer"
	Admin          = "admin"
	God            = "god"
)

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	School          string `json:"school"`
	CountryOfOrigin string `json:"countryOfOrigin"`
	YearOfStudy     int16  `json:"yearOfStudy"`
	Role            Role   `json:"role"`
}
