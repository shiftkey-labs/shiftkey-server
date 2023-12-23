package data

import "github.com/shiftkey-labs/shiftkey-server/pkg/model"

func GetLocalUsers() []model.User {
	users := []model.User{
		{
			ID:              "god",
			Name:            "Vansh Sood",
			Email:           "vanshsood@dal.ca",
			Role:            model.God,
			Interests:       []string{"Reading", "Swimming"},
			Major:           "All",
			School:          "none",
			CountryOfOrigin: "Moon",
			YearOfStudy:     1800,
			Age:             125,
		},
		{
			ID:              "u1",
			Name:            "Alice Johnson",
			Email:           "alice.johnson@example.com",
			Role:            model.Student,
			Interests:       []string{"Reading", "Swimming"},
			Major:           "Computer Science",
			School:          "University A",
			CountryOfOrigin: "USA",
			YearOfStudy:     2,
			Age:             20,
		},
		{
			ID:              "u2",
			Name:            "Bob Smith",
			Email:           "bob.smith@example.com",
			Role:            model.Volunteer,
			Interests:       []string{"Coding", "Music"},
			Major:           "Electrical Engineering",
			School:          "University B",
			CountryOfOrigin: "Canada",
			YearOfStudy:     3,
			Age:             21,
		},
	}

	return users
}

func GetLocalEvents() []model.Event {
	events := []model.Event{
		{
			ID:                    "e1",
			Title:                 "Tech Conference 2023",
			Image:                 "imageURL1",
			Description:           "Annual tech conference.",
			Location:              "Convention Center, City A",
			UnidentifiedAttendees: 10,
			Status:                model.Published,
			Date:                  "2023-07-10",
			Time:                  "10:00",
		},
		{
			ID:                    "e2",
			Title:                 "Science Fair",
			Image:                 "imageURL2",
			Description:           "Local science fair for high school students.",
			Location:              "High School B, City B",
			UnidentifiedAttendees: 5,
			Status:                model.Draft,
			Date:                  "2023-09-15",
			Time:                  "09:00",
		},
	}

	return events
}

func GetLocalHosts() []model.Host {
	hosts := []model.Host{
		{
			ID:          "h1",
			Name:        "Global Tech Inc.",
			Email:       "contact@globaltech.com",
			Company:     "Global Tech",
			Headshot:    "host1imageURL",
			Description: "Leading technology solutions provider.",
		},
		{
			ID:          "h2",
			Name:        "Innovate Green",
			Email:       "info@innovategreen.com",
			Company:     "Innovate Green",
			Headshot:    "host2imageURL",
			Description: "Environmental innovation organization.",
		},
	}

	return hosts
}

func GetLocalEventHosts() []model.EventHost {
	eventHosts := []model.EventHost{
		{
			ID:      "eh1",
			EventID: "e1",
			HostID:  "h1",
		},
		{
			ID:      "eh2",
			EventID: "e2",
			HostID:  "h2",
		},
	}

	return eventHosts
}

func GetLocalAttendances() []model.Attendance {
	attendances := []model.Attendance{
		{
			ID:      "a1",
			EventID: "e1",
			UserID:  "u1",
		},
		{
			ID:      "a2",
			EventID: "e2",
			UserID:  "u2",
		},
	}

	return attendances
}
