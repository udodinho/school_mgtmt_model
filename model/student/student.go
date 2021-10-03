package student

import (
	"school_mgmt/model/courses"
)

type Student struct {
	FirstName, LastName, Dept string
	Age, Score                int
	Course []string
}

var StudentProfile = map[int]Student{
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
	60:{ FirstName: "Kingsley",
		LastName: "Nnamani",
		Age:      25,
		Dept:     "ICT",
		Score:    65,
	},
	70:{ FirstName: "Emmanuel",
		LastName: "Njoku",
		Age:      22,
		Dept:     "Computer Science",
		Score:    40,
	},
}

var Victor = Student{
	"Victor",
	"Kelechi",
	"Technology",
	23,
	89,
	[]string{},
}

var course = courses.Courses

func (s *Student) TakeCourse(key int) map[int]Student {
	s.Course = append(s.Course, course[key])
	return nil
}