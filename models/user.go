package models

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

// always return error for something that can go wrong
// leave it upto caller to decide how its going to deal with that
func AddUser(u User) (User, error) {
	u.ID	= nextID
	nextID++
	users = append(users, &u)
	return u, nil
}


