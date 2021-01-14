package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// add/bind a method to a custom type
// userController automatically implements the Handler interface https://golang.org/pkg/net/http/#Handler
func (uc userController)  ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// convention for constructor function
// can return object itself, but return pointer to prevent unnecessary copy operation
func newUserController() *userController{
	// addressOf operator (&) is used on the variable but we don't declare the variable
	// this is allowed on struct data types
	return &userController{userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`)}
}
