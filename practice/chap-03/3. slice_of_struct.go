package main

import "fmt"

type Student struct {
	Roll int
	Name string
}

func main() {
	// students := []Student{}
	// for i := 1; i <= 10; i++ {
	// 	std := Student{Roll: i, Name: fmt.Sprintf("Student-%d", i)}
	// 	students = append(students, std)
	// }

	// // fmt.Println("Students list:", students)
	// for _, s := range students {
	// 	fmt.Println(s.Roll, "-->", s.Name)
	// }

	// // // Update students record
	// // for i, s := range students {
	// // 	fmt.Println("i:", i) // index
	// // 	fmt.Println("s:", s) // student object
	// // 	// students[i].Roll +=1 // It will work
	// // 	// s.Roll += 10 // Not update the original slice
	// // }

	// // Update students record
	// fmt.Println("Update all students roll no.")
	// for i := range students {
	// 	students[i].Roll += 10 // i--> index
	// }

	// fmt.Println("List Updated students details.")
	// for _, s := range students {
	// 	fmt.Println("Roll:", s.Roll, "-->", s.Name)
	// }

	// // Update students record
	// for _, s := range students {
	// 	s.Roll += 100
	// }

	// fmt.Println("Without Pointer Not Update students details.")
	// for _, s := range students {
	// 	fmt.Println("Roll:", s.Roll, "-->", s.Name)
	// }

	// ==================== With Pointer ================================
	students := []*Student{}
	for i := 1; i <= 10; i++ {
		std := Student{Roll: i, Name: fmt.Sprintf("Student-%d", i)}
		students = append(students, &std)
	}

	// Update students record
	for _, s := range students {
		s.Roll += 100
	}

	fmt.Println("Without Pointer Updated students details.")
	for _, s := range students {
		fmt.Println("Roll:", s.Roll, "-->", s.Name)
	}
}
