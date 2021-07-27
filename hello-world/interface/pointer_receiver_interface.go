package main

import "fmt"

type Career interface {
	career(istudy bool) string
}

type ComputerCareer struct {
	name        string
	application string
}

type SocialMedia struct {
	name        string
	application string
}

func (cc *ComputerCareer) career(istudy bool) string { // pointer receiver
	switch cc.name {
	case "coder":
		cc.name = "engineer"
	case "engineer":
		cc.name = "scientist"
	case "scientist":
		cc.name = "enterpriser"
	default:
		cc.name = "richer"
	}
	return cc.name
}

func (sm SocialMedia) career(istudy bool) string { // pointer receiver
	switch sm.name {
	case "war-journalist":
		sm.name = "journalist"
	case "writer":
		sm.name = "super-writer"
	case "scientist":
		sm.name = "enterpriser"
	default:
		sm.name = "richer"
	}
	return sm.name
}

func main() {
	var coder = ComputerCareer{"enterpriser", "blockchain"}
	var p Career
	var journalist = SocialMedia{"war-jounalist", "wars"}
	p = &coder
	var istudy = true
	fmt.Println(p.career(istudy))
	fmt.Println("######### coder's name is " + coder.name + "\n--------") // coder's name is changed
	p = journalist
	fmt.Println("######### journalist's name is " + journalist.name) // journalist's name is not changed
}
