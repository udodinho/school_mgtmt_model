package model

type Student struct {
	FirstName, LastName, Dept string
	Age, Score                int
	Course []string
}

var StudentProfile = map[int]Student  {
	10:{ FirstName: "Mercy",
		LastName: "Adebayo",
		Age:      24,
		Dept:     "MicroBiology",
		Score:    70,
	},
	20:{ FirstName: "John",
		LastName: "Okafor",
		Age:      19,
		Dept:     "Physics",
		Score:    39,

	},
	30:{ FirstName: "Khaleed",
		LastName: "Musa",
		Age:      23,
		Dept:     "Mathematics",
		Score:    59,
	},
	40:{ FirstName: "Olaoluwa",
		LastName: "Hammed",
		Age:      18,
		Dept:     "Chemical Engineering",

	},
	50:{ FirstName: "Faithia",
		LastName: "Olawale",
		Age:      20,
		Dept:     "FST",
		Score:    90,
	},