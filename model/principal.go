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

func (p Principal) Expulsion(key int) map[int]students.Student {
	for i := range StudentProfile {
		if key == i {
			delete(StudentProfile,i) 		}
	}
	return StudentProfile
}