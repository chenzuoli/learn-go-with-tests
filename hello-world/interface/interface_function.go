package main

import "fmt"

// interface defination
type SalaryCalculator interface {
	CalculateSalary() int
}

// employee1 defination
type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

// employee2 defination
type Contract struct {
	empId    int
	basicPay int
}

// implement the interface SalaryCalculator
func (permanent Permanent) CalculateSalary() int {
	return permanent.basicPay + permanent.pf
}

// implement the interface SalaryCalculator
func (contract Contract) CalculateSalary() int {
	return contract.basicPay
}

// calculate the total salary of the employee.
func totalSalary(employees []SalaryCalculator) int {
	var total int
	for _, employee := range employees {
		total += employee.CalculateSalary()
	}
	return total
}

func main() {
	// new some different type of employees
	var permantEmp Permanent = Permanent{
		empId:    1,
		basicPay: 15000,
		pf:       1000,
	}
	var contractEmp Contract = Contract{
		empId:    2,
		basicPay: 10000,
	}
	// calculate the total salary.
	fmt.Println(totalSalary([]SalaryCalculator{permantEmp, contractEmp}))
}
