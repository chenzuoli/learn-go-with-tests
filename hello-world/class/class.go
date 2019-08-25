package main

import employee "github.com/quii/learn-go-with-tests/hello-world/class/model"

func main() {
	var e1 = employee.New("chenzuoli", true, 1000000) // 避免了创建无用户名、性别、薪水的employee
	e1.Employee_information()
}
