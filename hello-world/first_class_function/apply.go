package main

import "fmt"

type student struct {
	firstname string
	lastname  string
	grade     string
	country   string
	height    int
}

func filterd(ss []student, f func(s student) bool) []student {
	var filterd_student []student
	for _, s_temp := range ss {
		if f(s_temp) {
			filterd_student = append(filterd_student, s_temp)
		}
	}
	return filterd_student
}
func doubled(ss []student, f func(s student) student) []student {
	var doubled_student []student
	for _, s_temp := range ss {
		s_temp = f(s_temp)
		doubled_student = append(doubled_student, s_temp)
	}
	return doubled_student
}

func main() {
	s1 := student{
		firstname: "chen",
		lastname:  "zuoli",
		grade:     "B",
		country:   "China",
		height:    170,
	}
	s2 := student{
		firstname: "wang",
		lastname:  "ying",
		grade:     "A",
		country:   "America",
		height:    160,
	}
	ss := []student{s1, s2}
	f1 := filterd(ss, func(s student) bool {
		if s.grade == "A" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(f1, "\n-----------------")
	f2 := doubled(ss, func(s student) student {
		s.height = s.height * 2
		return s
	})
	fmt.Println(f2)
}
