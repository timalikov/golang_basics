package main22

import "fmt"

type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee
	Department string
}

func (e Employee) Work() {
	fmt.Printf("Employee Name: %s, ID: %d is working.\n", e.Name, e.ID)
}

func main() {
	manager := Manager{
		Employee:   Employee{Name: "Tima", ID: 123},
		Department: "Sales",
	}

	manager.Work()

	fmt.Printf("Manager Department: %s\n", manager.Department)
}
