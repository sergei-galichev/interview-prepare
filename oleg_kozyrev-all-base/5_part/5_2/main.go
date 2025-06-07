package main

func main() {
	println(handle().Error())
}

type CustomError struct {
	message string
}

func (c *CustomError) Error() string {
	return c.message
}

func handle() error {
	return &CustomError{"this is a custom error"}
}
