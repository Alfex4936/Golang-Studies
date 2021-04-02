package main

func main() {
	government := NewGovern()
	A_university := &A_UNIV{
		mediator:  government,
		students:  20,
		fee:       300,
		isPrivate: false,
	}
	B_university := &B_UNIV{
		mediator:  government,
		students:  49,
		fee:       980,
		isPrivate: true,
	}

	A_university.MakePublic()
	B_university.MakePublic()

	A_university.GetMoreStudents(20)
	B_university.GetMoreStudents(50)

	A_university.IncreaseTuition(300)
	B_university.IncreaseTuition(10)
}
