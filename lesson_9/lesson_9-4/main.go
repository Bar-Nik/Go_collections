package main

import "fmt"

// 4. Создание вложенных структур: Создайте структуру "Employee" с полем
// "Name" и вложенной структурой "Job" с полями "Title" и "Salary".
// Создайте экземпляр этой структуры, заполните все поля и выведите их на
// экран.

type Job struct {
	Title  string
	Salary int
}
type Employee struct {
	Name string
	Job  Job
}

func main() {
	e := Employee{
		Name: "Li",
		Job: Job{
			Title:  "worker",
			Salary: 100000,
		},
	}
	fmt.Println(e)
}
