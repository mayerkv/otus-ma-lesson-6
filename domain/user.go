package domain

type User struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
}

func NewUser(ID int, username string, firstName string, lastName string, email string, phone string) *User {
	return &User{ID: ID, Username: username, FirstName: firstName, LastName: lastName, Email: email, Phone: phone}
}

func (u *User) ChangeEmail(email string) {
	u.Email = email
}

func (u *User) ChangePhone(phone string) {
	u.Phone = phone
}

func (u *User) ChangeLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) ChangeFirstName(firstName string) {
	u.FirstName = firstName
}
