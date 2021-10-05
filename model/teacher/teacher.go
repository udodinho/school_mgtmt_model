package teacher

import (
	"fmt"
	"school_mgmt/model/student"
)

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


func (t Teachers) Grading() string{
	for _, val := range student.StudentProfile {
		if val.Score >= 70 {
			fmt.Printf("Congrats %v you got an A in your examination\n", val)
		} else if val.Score >= 50 && val.Score <= 69 {
			fmt.Printf("Congrats %v you got a B in your examination\n", val)
		} else if val.Score >= 40 && val.Score <= 49 {
			fmt.Printf("Congrats %v you got a C in your examination\n", val)
		} else {
			fmt.Printf("%v You got an F in your examination\n", val)
		}
	}

	return ""
}