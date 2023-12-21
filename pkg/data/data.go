package data

import "github.com/shiftkey-labs/shiftkey-api/pkg/model"

func GetLocalUsers() []model.User {
	users := []model.User{
		{
			ID:              0,
			Name:            "vsood",
			School:          "NULL",
			CountryOfOrigin: "NULL",
			YearOfStudy:     1000,
			Role:            model.God,
		},
		{
			ID:              1,
			Name:            "Alice Johnson",
			School:          "University of Technology",
			CountryOfOrigin: "USA",
			YearOfStudy:     2,
			Role:            model.Student,
		},
		{
			ID:              2,
			Name:            "Bob Smith",
			School:          "Global Science College",
			CountryOfOrigin: "Canada",
			YearOfStudy:     3,
			Role:            model.Volunteer,
		},
		{
			ID:              3,
			Name:            "Carlos Garcia",
			School:          "Institute of Advanced Studies",
			CountryOfOrigin: "Spain",
			YearOfStudy:     1,
			Role:            model.Admin,
		},
		{
			ID:              4,
			Name:            "Diana Chen",
			School:          "National University",
			CountryOfOrigin: "China",
			YearOfStudy:     4,
			Role:            model.Student,
		},
	}

	return users
}
