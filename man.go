package main

import "fmt"

type Man struct {
	eyes int
	ears int
	legs int
}

func (man *Man) setEyes(noOfEyes int) {
	man.eyes = noOfEyes
}

func (man Man) getEyes() int {
	return man.eyes
}

func main() {
	var john Man
	john.setEyes(2)
	fmt.Println(john.getEyes())
}
