package main

import "fmt"

func main() {
	user := Action{
		Human{
			Name: "Grisha",
			Age:  23,
		},
	}
	user.Say("Hi, there")
}

type Action struct {
	Human
}

type Human struct {
	Name  string
	Age   int
	Email string
}

func (h Human) Say(msg string) {
	fmt.Printf("%s say: %s", h.Name, msg)
}
