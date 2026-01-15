package main

import "fmt"

type Address struct {
	village string
	state   string
}
type Students struct {
	rollno  int
	name    string
	subject string
	address Address
}

func NewStudent(rollno int, name, subject string) *Students {
	return &Students{
		rollno:  rollno,
		name:    name,
		subject: subject,
		// address: Address{
		// 	village: "Narotikapa",
		// 	state: "C.G",
		// },
	}
}

func UpdateStudent(rollno int, name, subject string, s *Students) {
	s.rollno = rollno
	s.name = name
	s.subject = subject
}

func (s *Students) Update(rollno int, name, subject string) {
	s.rollno = rollno
	s.name = name
	s.subject = subject
}

func main() {
	// ------------ Using normal variable --------------
	// s := Students{
	// 	rollno:  101,
	// 	name:    "Abhishek",
	// 	subject: "Hindi",
	// 	address: Address{village: "Pand", state: "C.G"},
	// }

	// fmt.Println("Student:", s)
	// fmt.Println("Student Address:", s.address)

	// Two way we can use for struct
	// use variable or pointer variable.

	// // ------------ Using struct pointer --------------
	// s1 := NewStudent(101, "Abhishek", "CS")
	// fmt.Println("Student-1:", s1)
	// fmt.Println("Student pointer:", *s1)
	// fmt.Println("---------- with * --------------")
	// fmt.Println("Student roll:", s1.rollno)
	// fmt.Println("Student name:", s1.name)
	// fmt.Println("Student subject:", s1.subject)

	// *s1.rollno -> *(s1.rollno) this is invalid because s1.rollno not a pointer & * only applicable for pointer
	// (*s1).rollno // not need this, GO will automatically do derefferencing.

	// Automatic pointer dereferencing for struct fields
	// s1.rollno = 200 | (*s1).rollno = 200
	// s1.rollno | (*s1).rollno
	// 	fmt.Println("Student roll with *:", (*s1).rollno)

	// 	But dereferencing IS required in other cases
	// Example with normal pointer
	// x := 10
	// p := &x

	// fmt.Println(*p)  // ✅ REQUIRED
	// fmt.Println(p)   // prints address

	// ---------------------------------------------------
	// ------------ Using struct pointer --------------
	s1 := NewStudent(101, "Abhishek", "CS")
	fmt.Println("Student-1:", s1)
	fmt.Println("Student pointer:", *s1)
	fmt.Println("------------------------")
	fmt.Println("Student roll:", s1.rollno)
	fmt.Println("Student name:", s1.name)
	fmt.Println("Student subject:", s1.subject)
	fmt.Println("----------- UpdateStudent()-------------")
	UpdateStudent(1, "Abhi", "GO", s1) // here: &s1 is wrong because s1 is already a pointer so (**Student wrong)
	fmt.Println("Student roll:", s1.rollno)
	fmt.Println("Student name:", s1.name)
	fmt.Println("Student subject:", s1.subject)
	fmt.Println("----------- Update()-------------")
	// Go takes the address automatically when calling pointer receiver methods
	// s1.Update(10, "Abhi Dahariya", "GO & Python")
	(s1).Update(10, "Abhi Dahariya", "GO & Python")
	// Go treats this -> s1.Update(...) to (&s1).Update(...).
	fmt.Println("Student roll:", s1.rollno)
	fmt.Println("Student name:", s1.name)
	fmt.Println("Student subject:", s1.subject)
}
