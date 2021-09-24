package model

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
	ID int
}

type Principal struct {
	Details Person
	Rank string
}

var StudentProfile = students.StudentProfile

func (p Principal) Expulsion(key int) map[int]students.Student {
	for i := range StudentProfile {
		if key == i {
			delete(StudentProfile,i) 		}
	}
	return StudentProfile
}

var applicantBio = applicants.ApplicantsBio

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