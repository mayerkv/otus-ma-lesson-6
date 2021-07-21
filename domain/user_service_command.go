package domain

type CreateUserCommand struct {
	Username  string `json:"username" binding:"required,alphanum"`
	FirstName string `json:"firstName" binding:"required,alpha"`
	LastName  string `json:"lastName" binding:"required,alpha"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,e164"`
}

type GetUserCommand struct {
	Id int `json:"id" binding:"required,number" uri:"id"`
}

type UpdateUserCommand struct {
	Id        int    `json:"id" binding:"required,number" uri:"id"`
	FirstName string `json:"firstName" binding:"alpha"`
	LastName  string `json:"lastName" binding:"alpha"`
	Email     string `json:"email" binding:"email"`
	Phone     string `json:"phone" binding:"e164"`
}

type DeleteUserCommand struct {
	Id int `json:"id" binding:"required,number" uri:"id"`
}
