package model

type Applicant struct {
	firstName string
	lastName string
	Age int
	gender string
}


var ApplicantsBio = map[int]Applicant {
	1: {
		firstName: "Eunice",
		lastName: "Nkemakolam",
		Age: 30,
		gender: "Female",
	},