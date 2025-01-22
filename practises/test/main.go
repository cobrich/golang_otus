package main

import(
	"fmt"

	"test/greeter"
	"test/tasks"
)

func main(){

	fmt.Println(greeter.Greet("Bekzat"))

	fmt.Println(tasks.TitleCase("My name is Bekzat my name is bekzat", "my bekzat"))
}