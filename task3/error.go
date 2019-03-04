package main

type MyError struct {
	msg string
}

func (error *MyError) Error() string {
	return error.msg
}
