package main

import "fmt"

type person struct {
	Fname string
	Lname string
}

func speak(f string, l string) {
	fmt.Println("Hello", f, l, "- Have a nice day")
}

func main() {
	p1 := person{
		"Adam",
		"Carter",
	}

	speak(p1.Fname, p1.Lname)
}
