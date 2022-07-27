package main

import ("fmt")

type People struct {
	name string
}

func (ppl People) greet() {
	fmt.Printf("Hello my name is %s", ppl.name)
}

func main() {
	ppl1 := People{}
	ppl1.name = "rakka"
	ppl1.greet()
}


