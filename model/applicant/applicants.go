package applicant

type Applicant struct {
	firstName string
	lastName string
	Age int
	gender string
}


var ApplicantsBio = map[int]Applicant{
	1: {
		firstName: "Eunice",
		lastName: "Nkemakolam",
		Age: 30,
		gender: "Female",
	},
	2: {
		firstName: "James",
		lastName: "Ibori",
		Age: 15,
		gender: "Male",
	},
	3: {
		firstName: "Maryam",
		lastName: "Balogun",
		Age: 30,
		gender: "Female",
	},
	4: {
		firstName: "Peter",
		lastName: "Friday",
		Age: 20,
		gender: "Male",
	},
}