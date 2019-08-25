package employee

import "fmt"

type employee struct {
	Name   string // the field name's first letter must be upper case, then other package can visit it.
	Gender bool
	Salary float64
}

func New(name string, gender bool, salary float64) employee { // 避免了创建无用户名、性别、薪水的employee
	employee := employee{Name: name, Gender: gender, Salary: salary}
	return employee
}

func (e employee) Employee_information() { // the function name's first letter must be upper case, then other package can visit it
	fmt.Println(e.Name, e.Gender, e.Salary)
}
