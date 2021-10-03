package teacher

import "school_mgmt/model/student"

type Teachers struct {
	FirstName string
	LastName  string
	Age int
	Identity string

}

var seniorLecturer = Teachers{
	FirstName: "Alex",
	LastName:  "Okocha",
	Age:       40,
	Identity:  "Lecturer",
}
