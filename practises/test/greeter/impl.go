package greeter



func Greet(name string) string{
	greeter := "Hello "

	if len(name) == 0{
		greeter = greeter + "world"
	}

	return greeter + name
}