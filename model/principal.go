package model

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