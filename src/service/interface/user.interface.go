package Interface

type StatusCode struct {
	Code    int
	Message string
}

type UserData struct {
	StatusCode
	Data []User
}

type User struct {
	Name string 
	Age  int 
}

