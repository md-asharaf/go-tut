package main

import "fmt"

type Student struct {
	name   string
	age    uint8
	height float32
}

type Teacher struct {
	fullname string
}

func main() {
	var student Student = Student{
		name:   "Asharaf",
		age:    23,
		height: 5.9,
	}
	fmt.Println("Student", student)
	student2 := Student{"rohit", 22, 6.4}
	fmt.Println("Student 2", student2)
	for i := range student.name {
		fmt.Printf("Character at index %d: %c\n", i, student.name[i])
	}
	teacher := Teacher{
		fullname: "John Doe",
	}
	fmt.Println("Teacher", teacher)

	var principal = struct { // not reusable
		name         string
		college_name string
	}{"Principal Name", "XYZ College"}
	fmt.Println("Principal", principal)
}
