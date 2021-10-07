package main

import (
	"fmt"
	"school_mgmt/model/applicant"
	"school_mgmt/model/principal"
	"school_mgmt/model/student"
	"school_mgmt/model/teacher"
)

var headAdmin  =  principal.Principal {
	Details: principal.Person{
		FirstName: "Kufre",
		LastName: "Ndudu",
		Age: 60,
		Identity:     "VC" ,
	},
	Rank: "Vice Chancellor",
}

var seniorLecturer = teacher.Teachers{
	FirstName: "Alex",
	LastName:  "Okocha",
	Age:       40,
	Identity:  "Lecturer",
}

var StudentProfile = student.StudentProfile

var applicantBio = applicant.ApplicantsBio


func main() {
	//headAdmin.Expulsion(70)

	//expelled := headAdmin.Expulsion(70)
	//fmt.Println(expelled)
	//admitted := headAdmin.Admin()
	//fmt.Println(admitted)
	//graded := seniorLecturer.Grading()
	//fmt.Println(graded)

	var victor = student.Victor
	victor.TakeCourse(3)
	victor.TakeCourse(1)
	fmt.Println(victor)

}
