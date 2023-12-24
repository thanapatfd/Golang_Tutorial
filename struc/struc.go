package main

import "fmt"

type employee struct {
	employeeID    string
	employeeName  string
	employeePhone string
}

func main() {
	// employee1 := employee{
	// 	employeeID:    "001",
	// 	employeeName:  "Ford1",
	// 	employeePhone: "00000000"}
	// employee2 := employee{
	// 	employeeID:    "002",
	// 	employeeName:  "Ford2",
	// 	employeePhone: "11111111"}
	// fmt.Println("employee1 = ", employee1)
	// fmt.Println("employee2 = ", employee2)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:    "101",
		employeeName:  "Test",
		employeePhone: "000000001",
	}
	employee2 := employee{
		employeeID:    "102",
		employeeName:  "Test2",
		employeePhone: "000000002",
	}
	employee3 := employee{
		employeeID:    "103",
		employeeName:  "Test3",
		employeePhone: "000000003",
	}
	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)
	fmt.Println("employeeList: ", employeeList)
}
