package principal

import (
	"fmt"
	"school_mgmt/model/applicant"
	"school_mgmt/model/student"
)

type Person struct {
	FirstName string
	LastName string
	Age int
	Identity string
}

type Principal struct {
	Details Person
	Rank    string
}
var StudentProfile = student.StudentProfile


func (p Principal) Expulsion(key int) map[int]student.Student {
	for i := range StudentProfile {
		if key == i {
			delete(StudentProfile,i) 		}
	}
	return StudentProfile
}

var applicantBio = applicant.ApplicantsBio

var headAdmin  =  Principal{
	Details: Person{
		FirstName: "Kufre",
		LastName: "Ndudu",
		Age: 60,
		Identity:     "VC" ,
	},
	Rank: "Vice Chancellor",
}

func (p Principal) Admin() string {
	returnString := ""
	for _, val := range applicantBio {
		if val.Age >= 18  {
			//fmt.Printf("Congratulations %v you have gain a provisional admission\n", val)
			returnString = fmt.Sprintf("Congratulations %v you have gain a provisional admission\n", val)
		}
	}
	return returnString
}